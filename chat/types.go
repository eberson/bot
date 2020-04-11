package chat

import (
	"regexp"
)

type Parameters map[string]string

type Config struct {
	Plugins  map[string]interface{}
	Entities []Entity
	Intents  []Intent
}

type Slack struct {
	Token string
	User  string
}

type GitHub struct {
	URL    string `yaml:"url"`
	APIURL string `yaml:"apiurl"`
	Token  string `yaml:"token"`
}

type Entity struct {
	Name         string
	Format       string
	Missing      string
	DefaultValue string
	Values       []string
}

type Intent struct {
	Expression    []string
	regex         []*regexp.Regexp
	Entities      []Entity
	NamedEntities []string
	Action        string
	Response      *Response
}

type Response struct {
	Template  string
	Messenger string
}

type ConversationState struct {
	Parameters Parameters
	Intent     Intent
	User       User
}

type User interface {
	Name() string
}

type ConversationStates map[string]ConversationState

type Conversation struct {
	intents []Intent
	context Context
	states  ConversationStates
}
