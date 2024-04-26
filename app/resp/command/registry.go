package command

import (
	"fmt"
	"strings"
)

type commandConstructor func() Command

var commandConstructors = map[string]commandConstructor{}

func GetCommand(name string) (Command, error) {
	if constructor, present := commandConstructors[strings.ToLower(name)]; present {
		return constructor(), nil
	} else {
		return nil, fmt.Errorf("command with name %v not found", name)
	}
}

func registerCommand(constructor commandConstructor) {
	c := constructor()

	commandConstructors[strings.ToLower(c.Name())] = constructor
}
