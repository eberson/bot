package slack

import (
	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
	"github.com/slack-go/slack"
)

type messenger struct {
	rtm *slack.RTM
}

func (m *messenger) Name() string {
	return PluginName
}

func (m *messenger) StartTyping(event events.Event) error {
	ev := event.Source().(*slack.MessageEvent)

	m.rtm.SendMessage(m.rtm.NewTypingMessage(ev.Channel))

	return nil
}

func (m *messenger) Send(event events.MessageEvent) error {
	ev := event.Event.Source().(*slack.MessageEvent)

	m.rtm.SendMessage(m.rtm.NewOutgoingMessage(event.Text(), ev.Channel))

	return nil
}

func newMessenger(rtm *slack.RTM) chat.Messenger {
	return &messenger{
		rtm: rtm,
	}
}
