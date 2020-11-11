package main

import (
	"github.com/tamoore/dada/internal/config"
	"github.com/tamoore/dada/internal/util"
)

func main() {
	c := make(chan config.Product)
	go config.Start(c)
	configProduct := <-c
	installStack := &util.ConcreteStack{}

	for _, program := range configProduct.Dependencies {
		installStack.Push(program)
	}
}
