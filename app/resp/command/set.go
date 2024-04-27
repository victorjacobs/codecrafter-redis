package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
)

type Set struct {
	key    string
	value  string
	expiry int64 // Unix nanos when this entry will expire
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

	i := 3
	for i < request.Len() {
		arg, _ := request.StringAt(i)
		if strings.ToLower(arg) == "px" {
			// Next array element is expiry, read it and move i beyond it
			i++

			expiryString, present := request.StringAt(i)
			if !present {
				return errors.New("PX argument missing param")
			}

			expiry, err := strconv.Atoi(expiryString)
			if err != nil {
				return fmt.Errorf("failed to parse argument as int: %w", err)
			}

			s.expiry = time.Now().UnixNano() + int64(expiry*1_000_000)
		}

		i++
	}

	return nil
}

func (s *Set) Key() string {
	return s.key
}

func (s *Set) Value() string {
	return s.value
}

func (s *Set) Expiry() int64 {
	return s.expiry
}

func init() {
	registerCommand(func() Command {
		return NewSet()
	})
}
