package main

import (
	"fmt"
	"os"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/plugins"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const cfgFileName = ".rootinha-bot.yaml"

var cfgFile string

func main() {
	cmd := newCmd()

	if err := cmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func newCmd() *cobra.Command {
	cobra.OnInitialize(initConfig)

	rootCmd := &cobra.Command{
		Use:   "rootinha-bot",
		Short: "",
		Long:  `.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := chat.CurrentConfig()

			context, err := plugins.NewContext(*config)

			if err != nil {
				return err
			}

			return chat.New(context).Start()
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s)", cfgFileName))

	return rootCmd
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = cfgFileName
	}

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Error("error reading the config file")
		return
	}

	logrus.WithField("filename", viper.ConfigFileUsed()).Info("using config file...", viper.ConfigFileUsed())
}
