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
	// viper.SetConfigName("config") // Looking for config.yaml
	// viper.SetConfigType("yaml")   // Expecting a YAML file
	// viper.AddConfigPath(".")      // Search in the current directory

	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal the config into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// log
	// fmt.Printf("%#v\n", config.Token)
	// fmt.Printf("%#v\n", config.Prefix)
	// fmt.Printf("%#v\n", config.Superusers)
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
