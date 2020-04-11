package console

import (
	"fmt"

	"github.com/eberson/rootinha/chat"

	"github.com/eberson/rootinha/events"
)

type messenger struct {
}

func NewMessenger() chat.Messenger {
	return &messenger{}
}

func (m *messenger) Name() string {
	return "console"
}

func (m *messenger) StartTyping(event events.Event) error {
	fmt.Println("bot started typing...")
	return nil
}

func (m *messenger) Send(event events.MessageEvent) error {
	fmt.Printf("%s\n", event.Text())
	return nil
}
