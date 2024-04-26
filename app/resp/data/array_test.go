package data

import (
	"bytes"
	"testing"
)

func TestUnmarshalSimpleArray(t *testing.T) {
	str := "*2\r\n+first\r\n+second\r\n"

	r := bytes.NewReader([]byte(str))

	array := NewArray().(*Array)
	err := array.UnmarshalBinary(r)
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if len(array.Elements()) != 2 {
		t.Errorf("array should contain 2 elements, contained %v", len(array.Elements()))
	} else if str, ok := array.Elements()[0].(*SimpleString); !ok {
		t.Errorf("could not convert first array element to SimpleString")
	} else if str.Data() != "first" {
		t.Errorf("expected first data element to contain 'first', contained %v", str.Data())
	}
}
