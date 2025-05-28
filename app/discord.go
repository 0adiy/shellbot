package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	CHAR_LIMIT = 1900
	MAX_CHUNKS = 3
)

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore anything that's not by a super user
	if !config.isSuperUser(m.Author.ID) {
		return
	}

	// log
	// fmt.Printf("Message Received: %v", m.Content)

	if strings.HasPrefix(m.Content, config.Prefix) {
		cmdString := strings.TrimSpace(strings.ToLower(m.Content))
		if cmdString == "?ping" {
			handlePing(s, m)
			return
		}

		handleShellExec(s, m)
	}

}

func sendLargeMessage(out string, s *discordgo.Session, m *discordgo.MessageCreate) {
	chunks := []string{}

	for len(out) > 0 && len(chunks) < MAX_CHUNKS {
		if len(out) > CHAR_LIMIT {
			chunks = append(chunks, out[:CHAR_LIMIT])
			out = out[CHAR_LIMIT:]
		} else {
			chunks = append(chunks, out)
			out = ""
		}
	}

	for i, chunk := range chunks {
		// If it's the last allowed chunk and there's still more text left, truncate and add "..."
		if i == MAX_CHUNKS-1 && len(out) > 0 {
			chunk = chunk[:len(chunk)-3] + "..."
		}
		s.ChannelMessageSend(m.ChannelID, "```ansi"+chunk+"```")
	}
}

func handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	apiLatency := time.Duration(s.HeartbeatLatency().Milliseconds()) * time.Millisecond

	// Calculate latency based on the difference between the received time and the sent message timestamp
	latency := time.Since(m.Timestamp)

	embed := &discordgo.MessageEmbed{
		Title: "Pong!",
		Color: 0x00FF00, // Green color for success
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Latency",
				Value:  fmt.Sprintf("%v", latency),
				Inline: true,
			},
			{
				Name:   "API Latency",
				Value:  fmt.Sprintf("%v", apiLatency),
				Inline: true,
			},
		},
	}

	// Send the embed as a response
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func handleShellExec(s *discordgo.Session, m *discordgo.MessageCreate) {
	command := strings.TrimSpace(m.Content[1:])

	shell := exec.Command("bash", "-c", command)

	out, err := shell.CombinedOutput()
	if err != nil {
		errStrToSend := fmt.Sprintf("Error : ```ansi\n%s\n```", err.Error())
		_, err = s.ChannelMessageSend(m.ChannelID, errStrToSend)
		if err != nil {
			fmt.Println("failed to send message", err.Error())
		}

		err = s.MessageReactionAdd(m.ChannelID, m.Message.ID, config.RejectedEmoji)
		if err != nil {
			fmt.Println("failed to add RejectedEmoji reaction", err.Error())
		}

		return
	}

	err = s.MessageReactionAdd(m.ChannelID, m.Message.ID, config.SuccessEmoji)
	if err != nil {
		fmt.Println("failed to add SuccessEmoji reaction", err.Error())
	}

	if string(out) != "" {
		sendLargeMessage(string(out), s, m)
	}
}
