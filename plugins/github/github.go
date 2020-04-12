package github

import (
	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/helper/strs"
	"github.com/pkg/errors"
)

const PluginName = "github"

type githubPlugin struct {
	actions chat.Actions
}

func (g *githubPlugin) Name() string {
	return PluginName
}

func (g *githubPlugin) Actions() chat.Actions {
	return g.actions.Actions()
}

func (g *githubPlugin) ActionByName(name string) chat.Action {
	return g.actions.ActionByName(name)
}

func New(config chat.Config) (chat.Plugin, error) {
	url, token, err := extractConfiguration(config)

	if err != nil {
		return nil, err
	}

	gh, err := NewGithub(url, token)

	if err != nil {
		return nil, err
	}

	actions := make(chat.Actions)
	actions.Add(
		newPullRequestsByState(gh),
	)

	return &githubPlugin{
		actions: actions,
	}, nil
}

func extractConfiguration(config chat.Config) (string, string, error) {
	conf, ok := config.Plugins["github"].(map[string]interface{})

	if !ok {
		return strs.Empty(), strs.Empty(), errors.New("we should receive a valid url and token for this plugin")
	}

	url, urlOK := conf["url"].(string)
	token, tokenOK := conf["token"].(string)

	if !urlOK || !tokenOK {
		return strs.Empty(), strs.Empty(), errors.New("we should receive a valid url and token for this plugin")
	}

	return url, token, nil
}

func Build(config chat.Config) chat.OptionFunc {
	return func(repository chat.Context) error {
		github, err := New(config)

		if err != nil {
			return err
		}

		repository.RegisterPlugin(github)
		return nil
	}
}
