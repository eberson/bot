package chat

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config Config

func CurrentConfig() *Config {
	if !structs.IsZero(config) {
		return &config
	}

	botConfig := viper.AllSettings()["bot"]

	if err := mapstructure.Decode(botConfig, &config); err != nil {
		logrus.WithError(err).Fatal("impossible to  read  config file")
	}

	if err := config.Validate(); err != nil {
		logrus.WithError(err).Fatal("error validating received configurations")
	}

	return &config
}

func (c *Config) Validate() error {
	validateIntents := func() error {
		for i, it := range c.Intents {
			if err := it.Validate(); err != nil {
				return err
			}

			c.Intents[i] = it
		}

		return nil
	}

	if err := validateIntents(); err != nil {
		return err
	}

	return nil
}
