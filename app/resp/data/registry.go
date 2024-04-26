package data

import "fmt"

type dataConstructor func() Data

var dataConstructors = map[string]dataConstructor{}

func getData(identifier string) (Data, error) {
	if constructor, present := dataConstructors[identifier]; present {
		return constructor(), nil
	} else {
		return nil, fmt.Errorf("command with name %v not found", identifier)
	}
}

func registerDataType(constructor dataConstructor) {
	c := constructor()

	dataConstructors[c.Identifier()] = constructor
}
