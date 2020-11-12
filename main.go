package main

import (
	"github.com/tamoore/dada/internal/config"
	"github.com/tamoore/dada/internal/setup"
	"github.com/tamoore/dada/internal/util"
)

func addNextInstallStep(queue *util.ConcreteQueue, packageManager string) *setup.InstallStep {
	stepName, hasNext := queue.Pop()
	step := &setup.InstallStep{}
	step.SetName(stepName.(string))
	step.SetPackageManager(packageManager)

	if hasNext {
		nextStep := addNextInstallStep(queue, packageManager)
		step.SetNext(nextStep)
	}

	return step
}

func main() {
	c := make(chan config.Product)
	go config.Start(c)
	configProduct := <-c
	installStack := &util.ConcreteQueue{}
	installMap := make(map[string]bool)

	for _, program := range configProduct.Dependencies {
		installMap[program] = false
		installStack.Push(program)
	}

	step := addNextInstallStep(installStack, configProduct.PackageManager)
	step.Execute(installMap)
}
