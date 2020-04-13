package slack

import "github.com/slack-go/slack"

type Slack struct {
	api          *slack.Client
	rootinha     slack.User
	rootinhaUser string
	users        []slack.User
	dms          []DirectMessage
}

func newSlack(api *slack.Client, user string) (*Slack, error) {
	client := &Slack{
		api:          api,
		rootinhaUser: user,
	}

	err := client.discoverUsers()

	if err != nil {
		return nil, err
	}

	return client, nil
}
