package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(NewModel())
	err := p.Start()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
