package github

import (
	"context"

	"github.com/fatih/structs"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
)

type listPullRequestByStateAction struct {
	github *Github

	owner      string
	repository string
	state      string
}

func newPullRequestsByState(github *Github) chat.Action {
	return &listPullRequestByStateAction{
		github: github,
	}
}

func (l *listPullRequestByStateAction) Name() string {
	return "list-pull-requests"
}

func (l *listPullRequestByStateAction) Fill(intent chat.Intent, parameters chat.Parameters) error {
	for _, entity := range intent.Entities {
		switch entity.Name {
		case "owner":
			if err := entity.ValueInto(parameters, &l.owner); err != nil {
				return err
			}
		case "repository":
			if err := entity.ValueInto(parameters, &l.repository); err != nil {
				return err
			}
		case "state":
			if err := entity.ValueInto(parameters, &l.state); err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *listPullRequestByStateAction) Run(intent chat.Intent, messenger chat.Messenger, event events.Event) error {
	prs, err := l.github.ListPullRequestsByState(context.Background(), l.owner, l.repository, l.state)

	if err != nil {
		return err
	}

	var pullRequests []interface{}

	for _, pr := range prs {
		pullRequests = append(pullRequests, structs.Map(pr))
	}

	return messenger.Send(events.MessageEvent{
		Event: event,
		Params: map[string]interface{}{
			"PRS": pullRequests,
		},
		Template: intent.Response.Template,
		To:       event.From(),
	})
}
