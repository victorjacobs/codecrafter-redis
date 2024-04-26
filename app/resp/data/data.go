package data

import "bytes"

type Data interface {
	Identifier() string
	// UnmarshalBinary unmarshals the Data from a reader, moving it beyond its bytes
	UnmarshalBinary(r *bytes.Reader) error
	MarshalBinary() ([]byte, error)
}
