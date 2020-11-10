package config

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Config ...
type Config struct {
	os     string
	config Product
}

// Init ...
func (c Config) Init() tea.Cmd {
	return nil
}

// Update ...
func (c Config) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	}
	return c, nil
}

// View ...
func (c Config) View() string {
	return fmt.Sprintf("Dada program")
}

// NewConfig ...
func NewConfig(config Product) *Config {
	return &Config{
		os:     config.PackageManager,
		config: config,
	}
}
