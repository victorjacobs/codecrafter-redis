package data

import (
	"bytes"
	"testing"
)

func TestUnmarshalSimpleString(t *testing.T) {
	str := "+1234\r\n"

	r := bytes.NewReader([]byte(str))

	simpleString := NewSimpleString().(*SimpleString)
	err := simpleString.UnmarshalBinary(r)
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if simpleString.Data() != "1234" {
		t.Errorf("expected string to contain 1234, contains %v", simpleString.Data())
	}
}
