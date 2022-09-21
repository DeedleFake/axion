package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Cancel key.Binding
	Size   tea.WindowSizeMsg
}

func NewModel() Model {
	return Model{
		Cancel: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.onKeyMsg(msg)
	case tea.WindowSizeMsg:
		m.Size = msg
	}

	return m, nil
}

func (m Model) onKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.Cancel):
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Width(m.Size.Width-2).
		Height(m.Size.Height-2).
		Border(lipgloss.RoundedBorder(), true).
		Align(.5, .5).
		Render("Alright.\nThat's pretty neat.")
}
