package main

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

func (s *Slack) loadBotUserID(api *slack.Client) error {
	users, err := api.GetUsers()

	if err != nil {
		return errors.Wrap(err, "error loading users from workspace")
	}

	for _, user := range users {
		if user.IsBot && strings.EqualFold(user.Name, s.User) {
			s.UserID = user.ID
			break
		}
	}

	logrus.WithFields(logrus.Fields{
		"user": s.User,
	}).Warn("it was not possible to load user id...")

	return nil
}
