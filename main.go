package main

import (
	"github.com/tamoore/dada/internal/config"
	"github.com/tamoore/dada/internal/setup"
	"github.com/tamoore/dada/internal/util"
)

func addNextInstallStep(stack *util.ConcreteStack, step setup.InstallStepable) {
	nextStepName, hasNext := stack.Pop()

	if hasNext {
		nextStep := &setup.InstallStep{}
		nextStep.SetName(nextStepName.(string))
		nextStep.SetPackageManager(step.GetPackageManager())
		addNextInstallStep(stack, nextStep)
	}
}

func main() {
	c := make(chan config.Product)
	go config.Start(c)
	configProduct := <-c
	installStack := &util.ConcreteStack{}

	for _, program := range configProduct.Dependencies {
		installStack.Push(program)
	}

	initialStepName, hasNext := installStack.Pop()
	if hasNext {
		initialStep := &setup.InstallStep{}
		initialStep.SetName(initialStepName.(string))
		initialStep.SetPackageManager(configProduct.PackageManager)
		addNextInstallStep(installStack, initialStep)
	}
}
