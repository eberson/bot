package console

import (
	"bufio"
	"os"
	"strings"

	"github.com/eberson/rootinha/helper/strs"

	"github.com/eberson/rootinha/chat"

	"github.com/sirupsen/logrus"

	"github.com/eberson/rootinha/events"
)

type input struct {
}

func (i *input) Name() string {
	return "console input"
}

func (i *input) Start() chan events.Event {
	c := make(chan events.Event)

	go func(channel chan events.Event) {
		reader := bufio.NewReader(os.Stdin)

		for {
			text, err := reader.ReadString('\n')

			if strs.IsEmpty(strings.Replace(text, "\n", strs.Empty(), -1)) {
				continue
			}

			if err != nil {
				logrus.WithError(err).Error("error reading data from console")
			}

			channel <- &event{
				text: text,
			}
		}
	}(c)

	return c
}

func NewInput() chat.Input {
	return &input{}
}

type event struct {
	text string
}

func (e event) Source() interface{} {
	return e.text
}

func (e event) From() string {
	return "nobody"
}

func (e event) Text() string {
	return e.text
}
