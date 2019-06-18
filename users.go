package main

import (
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DirectMessage struct {
	User    slack.User
	Channel string
}

type DirectMessages []DirectMessage

func (dms *DirectMessages) IsDMToBot(channel string) bool {
	for _, dm := range *dms {
		if dm.Channel == channel {
			return true
		}
	}

	return false
}

func loadDirectMessages(api *slack.Client) (DirectMessages, error) {
	users, err := api.GetUsers()

	if err != nil {
		return nil, errors.Wrap(err, "error loading users to load direct messages channels")
	}

	var directMessages DirectMessages

	for _, user := range users {
		if user.IsBot {
			continue
		}

		_, _, channel, err := api.OpenIMChannel(user.ID)

		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"userID":   user.ID,
				"username": user.Name,
			}).Error("error opening direct message channel")

			continue
		}

		directMessages = append(directMessages, DirectMessage{
			User:    user,
			Channel: channel,
		})
	}

	return directMessages, nil
}
