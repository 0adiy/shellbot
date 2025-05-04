# ShellBot - Discord VPS Shell Bot

ShellBot is a Discord bot written in Go (Golang) that lets you control your VPS directly from Discord. It provides a shell-like interface, allowing you to run commands on your VPS without the need to SSH directly. As of now, it only supports Linux systems (Bash).

## Features

- **Control your VPS through Discord**: Use Discord commands to interact with your VPS as though you're using a shell.
- **Superuser Support**: Only designated superusers can execute commands on the VPS.
- **Easy Setup**: Get up and running quickly with the `config.yaml` file for bot configuration and a `taskfile.yaml` for building the bot.

## Requirements

- **Go 1.18+** (for building the bot)
- **Linux-based VPS** (the bot currently only supports Bash shells)
- **Discord Bot Token** (you need to create a bot on the [Discord Developer Portal](https://discord.com/developers/applications))
- **Task** (optional but recommended for building and other tasks)

## Setup and Installation

### 1. Clone the repository

```bash
git clone https://github.com/0adiy/shellbot.git
cd shellbot
````

### 2. Create your `config.yaml`

Create a `config.yaml` file in `./bin/` dir (basically it needs to be in same folder as binary) of the repository with the following content:

```yaml
prefix: "?"  # The prefix to trigger the bot
token: "<Discord Bot Token>"  # Your Discord bot token
superusers:
  "user1": "757478713402064991"  # Replace with your Discord user IDs
  "user2": "829417226040901652"  # keys (names) here don't matter but can be used for username
  "user3": "418364415856082953"
```

* Replace `<Discord Bot Token>` with your actual Discord bot token.
* Replace the user IDs with the IDs of the users you want to designate as superusers (users who can execute commands).

### 3. Build the Bot

To build the project, make sure you have `Task` installed. Once that's set up, you can simply run:

```bash
task build
```

This will build the bot using the provided `taskfile.yaml`.

Alternatively build using bash 
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/shellbot ./app/.
```

### 4. Running the Bot

Once you've built the bot, make it executable 
```bash
chmod +x ./bin/shellbot
```

and then you can run it with the following command:

```bash
./bin/shellbot
```

### 5. Invite the Bot to Your Discord Server

1. Go to the [Discord Developer Portal](https://discord.com/developers/applications) and select your bot.
2. Under the "OAuth2" tab, generate an invite link for your bot with the `bot` permission.
3. Use the generated link to invite the bot to your server.

## Usage

Once the bot is running, you can start sending commands from your Discord server.

### Example Commands

Only superusers (listed in `config.yaml`) can run commands in shell, everyone else is ignored
* **Run a shell command**:

  Send a message like:

  ```txt
  ?ls -la
  ```

  This will run the `ls -la` command on your VPS and return the output in Discord.

## Contributing

Feel free to fork the repo, create issues, and send pull requests. If you find any bugs or have suggestions for features, let me know by creating an issue in the GitHub repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.