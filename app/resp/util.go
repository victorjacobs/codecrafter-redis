package resp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// ReadLine reads the reader until \r\n, skipping the delimiter
// and returning the data read.
func ReadLine(r *bytes.Reader) (string, error) {
	var buf bytes.Buffer

	for r.Len() > 0 {
		curChar, err := r.ReadByte()
		if err != nil {
			return "", err
		}

		buf.WriteByte(curChar)

		if buf.Len() >= 2 {
			b := buf.Bytes()
			if bytes.Equal(b[len(b)-2:], []byte("\r\n")) {
				return string(b[:len(b)-2]), nil
			}
		}
	}

	// If at this point we have not seen an \r\n the data is invalid
	return "", errors.New("data did not include trailing \\r\\n")
}

// PeekDataHeader peeks at the data header in reader. It does not move
// the reader and returns data header without delimiter
func PeekDataHeader(r *bytes.Reader) (string, error) {
	header, err := ReadLine(r)
	if err != nil {
		return "", fmt.Errorf("failed to read line: %w", err)
	}

	// Rewind to right before the header, which is len(line) + 2
	r.Seek(-int64(len(header)+2), io.SeekCurrent)

	return header, nil
}
