package data

import (
	"bytes"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func UnmarshalBinary(r *bytes.Reader) (Data, error) {
	header, err := resp.PeekDataHeader(r)
	if err != nil {
		return nil, err
	}

	data, err := getData(header[:1])
	if err != nil {
		return nil, err
	}

	err = data.UnmarshalBinary(r)
	if err != nil {
		return nil, err
	}

	return data, err
}
