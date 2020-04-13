package plugins

import (
	"errors"

	"github.com/eberson/rootinha/plugins/slack"

	"github.com/eberson/rootinha/plugins/console"

	"github.com/sirupsen/logrus"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/plugins/github"
)

type DefaultContext struct {
	inputs     map[string]chat.Input
	plugins    map[string]chat.Plugin
	messengers map[string]chat.Messenger
}

func NewContext(config chat.Config) (chat.Context, error) {
	var opts []chat.OptionFunc

	for name := range config.Plugins {
		switch name {
		case console.PluginName:
			opts = append(opts, console.Build(config))
		case github.PluginName:
			opts = append(opts, github.Build(config))
		case slack.PluginName:
			opts = append(opts, slack.Build(config))
		}
	}

	context := &DefaultContext{
		inputs:     make(map[string]chat.Input),
		plugins:    make(map[string]chat.Plugin),
		messengers: make(map[string]chat.Messenger),
	}

	for _, opt := range opts {
		if err := opt(context); err != nil {
			return nil, err
		}
	}

	return context, nil
}

func (c *DefaultContext) RegisterInput(input chat.Input) {
	if _, exists := c.plugins[input.Name()]; exists {
		logrus.WithFields(logrus.Fields{
			"name": input.Name(),
		}).Warn("input will be replaced")
	}

	c.inputs[input.Name()] = input
}

func (c *DefaultContext) RegisterPlugin(plugin chat.Plugin) {
	if _, exists := c.plugins[plugin.Name()]; exists {
		logrus.WithFields(logrus.Fields{
			"name": plugin.Name(),
		}).Warn("plugin will be replaced")
	}

	c.plugins[plugin.Name()] = plugin
}

func (c *DefaultContext) RegisterMessenger(messenger chat.Messenger) {
	if _, exists := c.messengers[messenger.Name()]; exists {
		logrus.WithFields(logrus.Fields{
			"name": messenger.Name(),
		}).Warn("messenger will be replaced")
	}

	c.messengers[messenger.Name()] = messenger
}

func (c *DefaultContext) Plugin(name string) (chat.Plugin, error) {
	plugin, exists := c.plugins[name]

	if !exists {
		return nil, errors.New("plugin not found")
	}

	return plugin, nil
}

func (c *DefaultContext) Messenger(name string) (chat.Messenger, error) {
	m, exists := c.messengers[name]

	if !exists {
		return nil, errors.New("there's no messenger registered for given name")
	}

	return m, nil
}

func (c *DefaultContext) Inputs() []chat.Input {
	var result []chat.Input

	for _, input := range c.inputs {
		result = append(result, input)
	}

	return result
}
