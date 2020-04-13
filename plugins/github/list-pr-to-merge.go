package github

import (
	"context"
	"strings"

	"github.com/fatih/structs"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
)

type listPullRequestsToMerge struct {
	github *Github

	owner      string
	repository string
	label      string
}

func newPullRequestsToMerge(github *Github) chat.Action {
	return &listPullRequestsToMerge{
		github: github,
	}
}

func (l *listPullRequestsToMerge) Name() string {
	return "list-pull-requests-to-merge"
}

func (l *listPullRequestsToMerge) Fill(intent chat.Intent, parameters chat.Parameters) error {
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
		case "label":
			if err := entity.ValueInto(parameters, &l.label); err != nil {
				return err
			}
		}

	}

	return nil
}

func (l *listPullRequestsToMerge) Run(intent chat.Intent, messenger chat.Messenger, event events.Event) error {
	prs, err := l.github.ListPullRequestsByState(context.Background(), l.owner, l.repository, "open")

	if err != nil {
		return err
	}

	var pullRequests []interface{}

	for _, pr := range prs {
		for _, label := range pr.Labels {
			if strings.EqualFold(*label.Name, l.label) {
				pullRequests = append(pullRequests, structs.Map(pr))
				break
			}
		}
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
