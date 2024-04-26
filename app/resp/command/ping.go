package command

import "github.com/codecrafters-io/redis-starter-go/app/resp/data"

type Ping struct{}

var _ Command = &Ping{}

func NewPing() Command {
	return &Ping{}
}

func (p *Ping) Name() string {
	return "PING"
}

func (p *Ping) UnmarshalData(request *data.Array) error {
	return nil
}

func init() {
	registerCommand(func() Command {
		return NewPing()
	})
}
