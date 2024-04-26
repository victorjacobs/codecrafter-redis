package command

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	str := "*2\r\n$4\r\necho\r\n$3\r\nhey\r\n"

	cmd, err := UnmarshalBinary([]byte(str))
	if err != nil {
		t.Errorf("err should be empty, was: %v", err)
	} else if echoCmd, isEcho := cmd.(*Echo); !isEcho {
		t.Errorf("expected ECHO command, got %T", cmd)
	} else if echoCmd.data != "hey" {
		t.Errorf("expected data to be 'hey', got %v", echoCmd.data)
	}
}
