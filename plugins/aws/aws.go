package aws

import (
	"github.com/eberson/rootinha/chat"
	"github.com/eberson/rootinha/helper/strs"
	"github.com/pkg/errors"
)

const PluginName = "aws"

type plugin struct {
	actions chat.Actions
}

func (p *plugin) Name() string {
	return PluginName
}

func (p *plugin) Actions() chat.Actions {
	return p.actions.Actions()
}

func (p *plugin) ActionByName(name string) chat.Action {
	return p.actions.ActionByName(name)
}

func New(config chat.Config) (chat.Plugin, error) {
	secretKey, accessKey, region, err := extractConfiguration(config)

	if err != nil {
		return nil, err
	}

	client, err := newAWS(secretKey, accessKey, region)

	if err != nil {
		return nil, err
	}

	actions := make(chat.Actions)
	actions.Add(
		newStartCodeBuild(client),
	)

	return &plugin{
		actions: actions,
	}, nil
}

func Build(config chat.Config) chat.OptionFunc {
	return func(repository chat.Context) error {
		aws, err := New(config)

		if err != nil {
			return err
		}

		repository.RegisterPlugin(aws)
		return nil
	}
}

func extractConfiguration(config chat.Config) (string, string, string, error) {
	conf, ok := config.Plugins["aws"].(map[string]interface{})

	if !ok {
		return strs.Empty(),
			strs.Empty(),
			strs.Empty(),
			errors.New("we should receive a valid secret key, access key and region for this plugin")
	}

	secretKey, secretKeyOK := conf["secret"].(string)
	accessKey, accessKeyOK := conf["access"].(string)
	region, regionOK := conf["region"].(string)

	if !secretKeyOK || !accessKeyOK || !regionOK {
		return strs.Empty(),
			strs.Empty(),
			strs.Empty(),
			errors.New("we should receive a valid secret key, access key and region for this plugin")
	}

	return secretKey, accessKey, region, nil
}
