package console

import (
	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
)

type greeting struct {
	name string
}

func newGreeting() chat.Action {
	return &greeting{}
}

func (g *greeting) Name() string {
	return "greeting"
}

func (g *greeting) Run(intent chat.Intent, messenger chat.Messenger, event events.Event) error {
	return messenger.Send(events.MessageEvent{
		Event:    event,
		Template: intent.Response.Template,
		To:       event.From(),
	})
}
