package aws

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
)

type startCodeBuild struct {
	client *AWS

	projectName string
}

func (s *startCodeBuild) Name() string {
	return "start-code-build"
}

func newStartCodeBuild(client *AWS) chat.Action {
	return &startCodeBuild{
		client: client,
	}
}

func (s *startCodeBuild) Fill(intent chat.Intent, parameters chat.Parameters) error {
	var filled int

	for _, entity := range intent.Entities {
		if strings.EqualFold(entity.Name, "project") {
			if err := entity.ValueInto(parameters, &s.projectName); err != nil {
				return err
			}

			filled++
			break
		}
	}

	if filled < 1 {
		return errors.New("there are fields that is not filled")
	}

	return nil
}

func (s *startCodeBuild) Run(intent chat.Intent, messenger chat.Messenger, event events.Event) error {
	id, err := s.client.build(s.projectName)

	if err != nil {
		return err
	}

	return messenger.Send(events.MessageEvent{
		Event: event,
		Params: map[string]interface{}{
			"BuildID": id,
		},
		Template: intent.Response.Template,
		To:       event.From(),
	})
}
