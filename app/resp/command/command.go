package command

import "github.com/codecrafters-io/redis-starter-go/app/resp/data"

type Command interface {
	Name() string
	UnmarshalData(request *data.Array) error
}
