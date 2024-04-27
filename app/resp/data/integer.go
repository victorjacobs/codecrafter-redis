package data

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Integer struct {
	data int
}

var _ Data = &Integer{}

func NewInteger() Data {
	return &Integer{}
}

func (i *Integer) Identifier() string {
	return ":"
}

func (i *Integer) MarshalBinary() ([]byte, error) {
	panic("unimplemented")
}

func (i *Integer) UnmarshalBinary(r *bytes.Reader) error {
	lineBytes, err := resp.ReadLine(r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal integer: %w", err)
	}

	int, err := strconv.Atoi(lineBytes[1:])
	if err != nil {
		return err
	}

	i.data = int

	return nil
}

func init() {
	registerDataType(func() Data {
		return NewInteger()
	})
}
