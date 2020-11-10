package config

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	cl "github.com/fatih/color"
	"github.com/tamoore/dada/internal/util"
)

// Config ...
type Config struct {
	os     string
	choice configType
	chosen bool
	config Product
}

// Init ...
func (c Config) Init() tea.Cmd {
	return nil
}

// Update ...
func (c Config) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			fallthrough
		case tea.KeyEsc:
			return c, tea.Quit
		}

		switch msg.String() {
		case "k", "up":
			c.choice--
			if c.choice < 0 {
				c.choice = 0
			}
		case "j", "down":
			c.choice++
			if c.choice > 1 {
				c.choice = 1
			}
		case "enter":
			c.chosen = true
			c.config = NewProduct(c.choice)
			return c, nil
		case "q":
			return c, tea.Quit
		}

	}

	return c, nil
}

// View ...
func (c Config) View() string {
	view := ""
	if !c.chosen {
		view += cl.New(cl.FgHiYellow).Sprint("==> Dada cofiguration")
		view += "\n"
		view += cl.New(cl.FgWhite).Sprint("Choose which OS this install should target?")
		view += "\n"
		view += util.Checkbox("OSX", c.choice == OSX)
		view += util.Checkbox("Linux", c.choice == Linux)
	}

	view += "\n"
	return view
}

// NewProduct ...
func NewProduct(t configType) Product {
	builder := GetBuilder(Linux)
	director := NewDirector(builder)
	return director.MakeConfig()
}

// Start ...
func Start() {
	p := tea.NewProgram(&Config{})
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
