package console

import (
	"github.com/eberson/rootinha/chat"
)

const PluginName = "console"

type plugin struct {
	actions chat.Actions
}

func (p *plugin) Name() string {
	return PluginName
}

func (p *plugin) Actions() chat.Actions {
	return p.actions.Actions()
}

func (p *plugin) ActionByName(name string) chat.Action {
	return p.actions.ActionByName(name)
}

func New(config chat.Config) chat.Plugin {
	actions := make(chat.Actions)
	actions.Add(
		newGreeting(),
		newTime(),
	)

	return &plugin{
		actions: actions,
	}
}

func Build(config chat.Config) chat.OptionFunc {
	return func(repository chat.Context) error {
		repository.RegisterPlugin(New(config))
		repository.RegisterInput(NewInput())
		repository.RegisterMessenger(NewMessenger())
		return nil
	}
}
