package chat

import (
	"fmt"

	"github.com/eberson/rootinha/helper/strs"
)

func NewMissingEntityError(entity Entity, text string) MissingEntityError {
	return MissingEntityError{
		entity: entity,
		text:   text,
	}
}

type MissingEntityError struct {
	entity Entity
	text   string
}

func (m MissingEntityError) HasMissingSet() bool {
	return strs.IsNotEmpty(m.entity.Missing)
}

func (m MissingEntityError) MissingQuestion() string {
	return m.entity.Missing
}

func (m MissingEntityError) Error() string {
	return fmt.Sprintf("parameter %s not found in given text: %s", m.entity.Name, m.text)
}
