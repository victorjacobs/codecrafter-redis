package data

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type SimpleString struct {
	data string
}

var _ Data = &SimpleString{}

func NewSimpleString() Data {
	return &SimpleString{}
}

func NewSimpleStringWithData(data string) Data {
	return &SimpleString{
		data: data,
	}
}

func (s *SimpleString) Identifier() string {
	return "+"
}

func (s *SimpleString) UnmarshalBinary(r *bytes.Reader) error {
	lineBytes, err := resp.ReadLine(r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal simple string: %w", err)
	}

	s.data = string(lineBytes[1:])

	return nil
}

func (s *SimpleString) MarshalBinary() ([]byte, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%v%v\r\n", s.Identifier(), s.Data()))

	return []byte(sb.String()), nil
}

func (s *SimpleString) Data() string {
	return s.data
}

func init() {
	registerDataType(func() Data {
		return NewSimpleString()
	})
}
