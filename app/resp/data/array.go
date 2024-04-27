package data

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Array struct {
	elements []Data
}

func NewArray() Data {
	return &Array{
		elements: []Data{},
	}
}

func (d *Array) Identifier() string {
	return "*"
}

func (d *Array) UnmarshalBinary(r *bytes.Reader) error {
	// Read the data header
	header, err := resp.ReadLine(r)
	if err != nil {
		return fmt.Errorf("failed to read data header: %w", err)
	}

	expectedArrayElements, err := strconv.Atoi(header[1:])
	if err != nil {
		return fmt.Errorf("failed to read number of array elements: %w", err)
	}

	for r.Len() > 0 && len(d.elements) != expectedArrayElements {
		element, err := UnmarshalBinary(r)
		if err != nil {
			return err
		}

		d.elements = append(d.elements, element)
	}

	return nil
}

func (d *Array) MarshalBinary() ([]byte, error) {
	return nil, nil
}

func (d *Array) Elements() []Data {
	return d.elements
}

func (d *Array) Len() int {
	return len(d.elements)
}

func (d *Array) StringAt(pos int) (string, bool) {
	if pos > len(d.Elements())-1 {
		return "", false
	}

	el := d.Elements()[pos]

	if simpleString, isSimpleString := el.(*SimpleString); isSimpleString {
		return simpleString.Data(), true
	} else if bulkString, isBulkString := el.(*BulkString); isBulkString {
		return bulkString.Data(), true
	}

	return "", false
}

func init() {
	registerDataType(func() Data {
		return NewArray()
	})
}
