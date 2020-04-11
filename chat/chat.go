package chat

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Rootinha struct {
	context Context
}

func New(ctx Context) *Rootinha {
	return &Rootinha{
		context: ctx,
	}
}

func (r *Rootinha) Start() error {
	inputsSize := len(r.context.Inputs())

	if inputsSize == 0 {
		return errors.New("there are no input to receive action, so I have no reason to live...")
	}

	var wg sync.WaitGroup

	wg.Add(inputsSize)

	for _, input := range r.context.Inputs() {

		go func(in Input, group *sync.WaitGroup) {
			defer group.Done()

			if err := r.prepareListener()(in); err != nil {
				logrus.WithFields(logrus.Fields{
					"error": err,
					"name":  in.Name(),
				}).Error("error in input that was responsible for get input events")
			}
		}(input, &wg)
	}

	wg.Wait()

	return nil
}

func (r *Rootinha) prepareListener() func(Input) error {
	return func(input Input) error {
		logrus.WithFields(logrus.Fields{
			"name": input.Name(),
		}).Info("starting a worker for receive input events")

		if validator, ok := input.(Validator); ok {
			if err := validator.Validate(*CurrentConfig()); err != nil {
				return errors.Wrap(
					err,
					"there are invalid or no enough information for validate input.. stopping it",
				)
			}
		}

		config := CurrentConfig()

		channelEvents := input.Start()

		for {
			select {
			case event, valid := <-channelEvents:
				if !valid {
					return nil
				}

				err := NewConversation(r.context, config.Intents...).Execute(event)

				if err != nil {
					logrus.WithError(err).Error("error executing conversation")
				}
			default:
				continue
			}
		}
	}
}
