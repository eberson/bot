package events

import (
	"bytes"
	"text/template"

	"github.com/eberson/rootinha/helper/strs"

	"github.com/sirupsen/logrus"
)

type Event interface {
	Source() interface{}
	From() string
	Text() string
}

type MessageEvent struct {
	Event    Event
	Params   map[string]interface{}
	Value    string
	Template string
	To       string
}

func (e *MessageEvent) Text() string {
	params := e.Params

	if params == nil {
		params = make(map[string]interface{})
	}

	params["To"] = e.To

	text, err := makeMessageText(e.Template, params)

	if err != nil {
		return err.Error()
	}

	return text
}

func makeMessageText(tmpl string, params map[string]interface{}) (string, error) {
	tmp, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err,
			"template": tmpl,
		}).Error("error parsing template of intent")

		return strs.Empty(), err
	}

	var tpl bytes.Buffer

	err = tmp.Execute(&tpl, params)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err,
			"template": tmpl,
			"params":   params,
		}).Error(err)

		return strs.Empty(), err
	}

	return tpl.String(), nil
}
