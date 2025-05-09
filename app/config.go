package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Prefix        string            `mapstructure:"prefix"`
	Token         string            `mapstructure:"token"`
	Superusers    map[string]string `mapstructure:"superusers"` // username -> id
	SuccessEmoji  string            `mapstructure:"successEmoji"`
	RejectedEmoji string            `mapstructure:"rejectedEmoji"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal the config into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.SuccessEmoji == "" {
		config.SuccessEmoji = "✅"
	}
	if config.RejectedEmoji == "" {
		config.RejectedEmoji = "❌"
	}
	return &config, nil
}

func (c *Config) isSuperUser(userID string) bool {
	for _, id := range c.Superusers {
		if id == userID {
			return true
		}
	}

	return false
}
