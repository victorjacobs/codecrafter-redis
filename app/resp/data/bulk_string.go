package data

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type BulkString struct {
	data string
	null bool
}

var _ Data = &BulkString{}

func NewBulkString() Data {
	return &BulkString{}
}

func NewNullBulkString() Data {
	return &BulkString{
		null: true,
	}
}

func NewBulkStringWithData(data string) Data {
	return &BulkString{
		data: data,
	}
}

func (b *BulkString) Identifier() string {
	return "$"
}

func (b *BulkString) UnmarshalBinary(r *bytes.Reader) error {
	_, err := resp.ReadLine(r)
	if err != nil {
		return fmt.Errorf("failed to skip data header: %w", err)
	}

	data, err := resp.ReadLine(r)
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

	b.data = data

	return nil
}

func (b *BulkString) MarshalBinary() ([]byte, error) {
	var sb strings.Builder

	if b.null {
		sb.WriteString("$-1\r\n")
	} else {
		sb.WriteString(fmt.Sprintf("$%d\r\n%v\r\n", len(b.data), b.data))
	}

	return []byte(sb.String()), nil
}

func (b *BulkString) Data() string {
	return b.data
}

func init() {
	registerDataType(func() Data {
		return NewBulkString()
	})
}
