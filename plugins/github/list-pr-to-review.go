package github

import (
	"context"
	"strings"

	"github.com/fatih/structs"

	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/events"
)

type listPullRequestsToReview struct {
	github *Github

	owner      string
	repository string
	label      string
}

func newPullRequestsToReview(github *Github) chat.Action {
	return &listPullRequestsToReview{
		github: github,
	}
}

func (l *listPullRequestsToReview) Name() string {
	return "list-pull-requests-to-review"
}

func (l *listPullRequestsToReview) Fill(intent chat.Intent, parameters chat.Parameters) error {
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

func (l *listPullRequestsToReview) Run(intent chat.Intent, messenger chat.Messenger, event events.Event) error {
	prs, err := l.github.ListPullRequestsByState(context.Background(), l.owner, l.repository, "open")

	if err != nil {
		return err
	}

	var pullRequests []interface{}

	for _, pr := range prs {
		needReview := true

		for _, label := range pr.Labels {
			if strings.EqualFold(*label.Name, l.label) {
				needReview = false
				break
			}
		}

		if needReview {
			pullRequests = append(pullRequests, structs.Map(pr))
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
