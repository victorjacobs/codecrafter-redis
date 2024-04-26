package resp

import (
	"bytes"
	"testing"
)

func TestReadLine(t *testing.T) {
	str := "abc\r\n4242\r\n"

	r := bytes.NewReader([]byte(str))

	line, err := ReadLine(r)
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if line != "abc" {
		t.Errorf("first line should be equal to abc, is %v", string(line))
	}

	line, err = ReadLine(r)
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if line != "4242" {
		t.Errorf("second line should be equal to 4242, is %v", string(line))
	}
}

func TestInvalidData(t *testing.T) {
	str := "abc"

	r := bytes.NewReader([]byte(str))

	_, err := ReadLine(r)
	if err == nil {
		t.Error("should've returned error, returned nothing")
	}
}

func TestPeekDataHeader(t *testing.T) {
	str := "abc\r\n4242\r\n"

	r := bytes.NewReader([]byte(str))

	dataHeader, err := PeekDataHeader(r)
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if dataHeader != "abc" {
		t.Errorf("data header should be 123, is %v", dataHeader)
	}

	dataHeader, err = PeekDataHeader(r)
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if dataHeader != "abc" {
		t.Errorf("data header should be 123, is %v", dataHeader)
	}
}
