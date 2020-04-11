package chat

import (
	"strings"

	"github.com/eberson/rootinha/events"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	errIntentNotFound = errors.New("intent not found for given text")
)

func NewConversation(context Context, intents ...Intent) *Conversation {
	return &Conversation{
		intents: intents,
		context: context,
	}
}

func (c *Conversation) Execute(event events.Event) error {
	intent, err := c.findIntent(event)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"text":  event.Text(),
		}).Error(err)
		return err
	}

	action, err := c.action(intent.Action)

	if err != nil {
		logrus.WithError(err).Error("error finding action")
		return err
	}

	params := intent.Parameters(event.Text())

	messenger, err := c.context.Messenger(intent.Response.Messenger)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"messenger": intent.Response.Messenger,
			"error":     err,
		}).Error(err)

		return err
	}

	if configurable, ok := action.(ConfigurableAction); ok {
		if err = configurable.Fill(*intent, params); err != nil {
			return c.handleError(messenger, event, err)
		}
	}

	if err = action.Run(*intent, messenger, event); err != nil {
		return c.handleError(messenger, event, err)
	}

	return nil
}

func (c *Conversation) action(name string) (Action, error) {
	data := strings.Split(name, ":")

	if len(data) != 2 {
		return nil, errors.Wrap(errors.New("invalid action name"), name)
	}

	plugin, err := c.context.Plugin(data[0])

	if err != nil {
		return nil, errors.Wrap(err, "error finding plugin to get action for conversation")
	}

	action := plugin.ActionByName(data[1])

	if action == nil {
		return nil, errors.New("it was not possible to find action for conversation")
	}

	return action, nil
}

func (c *Conversation) findIntent(event events.Event) (*Intent, error) {
	for _, it := range c.intents {
		if it.Matches(event.Text()) {
			return &it, nil
		}
	}

	return nil, errIntentNotFound
}

func (c *Conversation) handleError(messenger Messenger, event events.Event, err error) error {
	if me, ok := err.(MissingEntityError); ok {
		err = messenger.Send(events.MessageEvent{
			Event: event,
			Value: me.MissingQuestion(),
			To:    event.From(),
		})

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"missingError": me,
				"error":        err,
			}).Error("error trying to send missing question")

			return err
		}

		return nil
	}

	return err
}
