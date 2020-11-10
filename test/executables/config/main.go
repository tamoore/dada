package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tamoore/dada/internal/config"
)

func main() {
	builder := config.GetBuilder(config.Linux)
	director := config.NewDirector(builder)
	configProduct := director.MakeConfig()
	model := config.NewConfig(configProduct)

	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
