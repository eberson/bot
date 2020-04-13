package slack

import (
	"fmt"
	"strings"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
	"github.com/eberson/rootinha/helper/strs"
	"github.com/fatih/structs"
	"github.com/slack-go/slack"

	"github.com/sirupsen/logrus"
)

type input struct {
	rtm   *slack.RTM
	slack *Slack
}

func (i *input) Name() string {
	return PluginName
}

func (i *input) Start() chan events.Event {
	c := make(chan events.Event)

	go func(s *Slack, rtm *slack.RTM, channel chan events.Event) {
		for {
			select {
			case msg := <-rtm.IncomingEvents:
				switch ev := msg.Data.(type) {
				case *slack.MessageEvent:
					e := &slackEvent{
						source: ev,
						client: s,
					}

					if s.ShouldAnswer(e) {
						channel <- e
					}
				case *slack.RTMError:
					logrus.Error(ev.Error())
				case *slack.InvalidAuthEvent:
					logrus.Error("invalid credentials")
				default:
					logrus.Debug("nothing to say...")
					//Take no action
				}
			}
		}
	}(i.slack, i.rtm, c)

	return c
}

func newInput(rtm *slack.RTM, s *Slack) (chat.Input, error) {
	return &input{
		rtm:   rtm,
		slack: s,
	}, nil
}

type slackEvent struct {
	source *slack.MessageEvent
	client *Slack
	text   string
}

func (e *slackEvent) Source() interface{} {
	return e.source
}

func (e *slackEvent) From() string {
	var user slack.User

	for _, usr := range e.client.users {
		if strings.EqualFold(usr.ID, e.source.User) {
			user = usr
			break
		}
	}

	if !structs.IsZero(user) {
		if e.client.isDirectToBot(e.source.Channel) {
			return user.RealName
		}

		return fmt.Sprintf("<@%s>", user.ID)
	}

	return e.source.User
}

func (e *slackEvent) Text() string {
	if strs.IsEmpty(e.text) {
		e.text = e.source.Text
	}

	return e.text
}
