package chat

import "github.com/eberson/rootinha/events"

type Messenger interface {
	Name() string
	StartTyping(event events.Event) error
	Send(event events.MessageEvent) error
}

type ConfigurableAction interface {
	Fill(parameters Parameters) error
}

type Action interface {
	Name() string
	Run(intent Intent, messenger Messenger, event events.Event) error
}

type Actions map[string]Action

func (a *Actions) Add(actions ...Action) {
	for _, action := range actions {
		(*a)[action.Name()] = action
	}
}

type Plugin interface {
	Name() string
	Actions() Actions
	ActionByName(name string) Action
}

type Input interface {
	Name() string
	Start() chan events.Event
}

type Validator interface {
	Validate(config Config) error
}
