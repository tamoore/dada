package setup

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	cl "github.com/fatih/color"
)

// Step ...
type Step interface {
	Execute(steps map[string]bool)
	SetNext(step Step)
}

// InstallStepable ...
type InstallStepable interface {
	Step
	SetName(name string)
	SetPackageManager(name string)
	GetPackageManager() string
}

// InstallStep ...
type InstallStep struct {
	packageManager string
	program        string
	next           Step
}

// Execute ...
func (s *InstallStep) Execute(steps map[string]bool) {
	if steps[s.program] {
		fmt.Printf("\nStep already %s done.", s.program)
		s.next.Execute(steps)
		return
	}
	cl.New(cl.FgHiYellow).Printf("==> Running '%s install %s'\n", s.packageManager, s.program)
	cmd := exec.Command("sudo", s.packageManager, "install", "-y", s.program)
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatal(cl.New(cl.FgHiRed).Sprint(err))
	}

	steps[s.program] = true

	if s.next != nil {
		s.next.Execute(steps)
	}
}

// SetNext ...
func (s *InstallStep) SetNext(step Step) {
	s.next = step
}

// SetName ...
func (s *InstallStep) SetName(name string) {
	s.program = name
}

// SetPackageManager ...
func (s *InstallStep) SetPackageManager(name string) {
	s.packageManager = name
}

// GetPackageManager ...
func (s *InstallStep) GetPackageManager() string {
	return s.packageManager
}
