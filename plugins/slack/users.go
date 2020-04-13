package slack

import (
	"strings"

	"github.com/slack-go/slack"
)

type Users []slack.User

func (s *Slack) discoverUsers() error {
	users, err := s.api.GetUsers()

	if err != nil {
		return err
	}

	var contacts []slack.User

	for _, user := range users {
		if user.IsBot {
			if strings.EqualFold(user.Name, s.rootinhaUser) {
				s.rootinha = user
			}

			continue
		}

		contacts = append(contacts, user)
	}

	s.users = contacts
	return nil
}
