package command

import (
	"errors"

	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
)

type Set struct {
	key   string
	value string
}

var _ Command = &Set{}

func NewSet() Command {
	return &Set{}
}

func (s *Set) Name() string {
	return "SET"
}

func (s *Set) UnmarshalData(request *data.Array) error {
	if key, present := request.StringAt(1); present {
		s.key = key
	} else {
		return errors.New("key not found in command")
	}

	if value, present := request.StringAt(2); present {
		s.value = value
	} else {
		return errors.New("value not found in command")
	}

	return nil
}

func (s *Set) Key() string {
	return s.key
}

func (s *Set) Value() string {
	return s.value
}

func init() {
	registerCommand(func() Command {
		return NewSet()
	})
}
