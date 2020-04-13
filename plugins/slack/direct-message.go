package slack

import (
	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type DirectMessage struct {
	user    slack.User
	channel string
}

func (u *Users) createDMChannel(api *slack.Client) []DirectMessage {
	var dms []DirectMessage

	for _, user := range *u {
		_, _, channel, err := api.OpenIMChannel(user.ID)

		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"userID":   user.ID,
				"username": user.Name,
			}).Error("error opening direct message channel")

			continue
		}

		dms = append(dms, DirectMessage{
			user:    user,
			channel: channel,
		})
	}

	return dms
}
