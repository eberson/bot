package chat

type OptionFunc func(context Context) error

type Context interface {
	RegisterPlugin(plugin Plugin)
	RegisterMessenger(messenger Messenger)
	RegisterInput(input Input)
	Inputs() []Input
	Plugin(name string) (Plugin, error)
	Messenger(name string) (Messenger, error)
}
