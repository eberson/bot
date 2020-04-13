package slack

import (
	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/helper/strs"
	"github.com/slack-go/slack"

	"github.com/pkg/errors"
)

const PluginName = "slack"

type plugin struct {
	api *slack.Client
}

func (p *plugin) Name() string {
	return PluginName
}

func (p *plugin) Actions() chat.Actions {
	return chat.Actions{}
}

func (p *plugin) ActionByName(name string) chat.Action {
	return nil
}

func New(api *slack.Client) chat.Plugin {
	return &plugin{
		api: api,
	}
}

func Build(config chat.Config) chat.OptionFunc {
	return func(repository chat.Context) error {
		user, token, err := extractConfiguration(config)

		if err != nil {
			return err
		}

		api := slack.New(token, slack.OptionDebug(true))

		rtm := api.NewRTM()
		go rtm.ManageConnection()

		client, err := newSlack(api, user)

		if err != nil {
			return err
		}

		in, err := newInput(rtm, client)

		repository.RegisterPlugin(New(api))
		repository.RegisterInput(in)
		repository.RegisterMessenger(newMessenger(rtm))
		return nil
	}
}

func extractConfiguration(config chat.Config) (string, string, error) {
	conf, ok := config.Plugins["slack"].(map[string]interface{})

	if !ok {
		return strs.Empty(),
			strs.Empty(),
			errors.New("we should receive a valid token and bot user for this plugin")
	}

	user, userOK := conf["user"].(string)
	token, tokenOK := conf["token"].(string)

	if !userOK || !tokenOK {
		return strs.Empty(),
			strs.Empty(),
			errors.New("we should receive a valid token and bot user for this plugin")
	}

	return user, token, nil
}
