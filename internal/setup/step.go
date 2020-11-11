package setup

import "fmt"

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
	fmt.Printf("Run %s\n", s.program)
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
