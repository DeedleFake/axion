package main

import (
	"deedles.dev/axion/buffer"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Cancel key.Binding

	Size tea.WindowSizeMsg

	Buffer    buffer.Buffer
	Line, Col int
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

	c := m.Buffer.Cursor(m.Line, m.Col)
	switch msg.Type {
	case tea.KeyRunes, tea.KeySpace:
		c.Write([]byte(string(msg.Runes)))
		m.Line, m.Col = c.LineAndCol()
	case tea.KeyEnter:
		m.Buffer.Cursor(m.Line, m.Col).WriteByte('\n')
		m.Line, m.Col = c.LineAndCol()
	case tea.KeyTab:
		m.Buffer.Cursor(m.Line, m.Col).WriteByte('\t')
		m.Line, m.Col = c.LineAndCol()
	}

	return m, nil
}

func (m Model) View() string {
	if (m.Size.Width == 0) || (m.Size.Height == 0) {
		return ""
	}

	return lipgloss.NewStyle().
		Width(m.Size.Width-2).
		Height(m.Size.Height-2).
		Border(lipgloss.RoundedBorder(), true).
		Align(.5, .5).
		Render(string(m.Buffer.View(0, m.Size.Height-4)))
}
