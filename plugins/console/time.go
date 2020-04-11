package console

import (
	"time"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
)

type timeAction struct {
}

func newTime() chat.Action {
	return &timeAction{}
}

func (t *timeAction) Name() string {
	return "time"
}

func (t *timeAction) Run(intent chat.Intent, messenger chat.Messenger, event events.Event) error {
	return messenger.Send(events.MessageEvent{
		Event:    event,
		Template: intent.Response.Template,
		Params: map[string]string{
			"Time": time.Now().Format(time.RFC3339),
		},
		To: event.From(),
	})
}
