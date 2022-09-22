package main

import (
	"strings"

	"deedles.dev/axion/buffer"
	"deedles.dev/axion/theme"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var winBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "□",
	TopRight:    "┐",
	BottomLeft:  "└",
	BottomRight: "┘",
}

type Model struct {
	Cancel key.Binding

	Size        tea.WindowSizeMsg
	WinStyle    lipgloss.Style
	CursorStyle lipgloss.Style

	Buffer  buffer.Buffer
	BufView *buffer.View
	Cursor  *buffer.Cursor
}

func NewModel() Model {
	m := Model{
		Cancel: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),

		WinStyle: theme.Default().Editor().
			Border(winBorder, true),
		CursorStyle: theme.Default().Editor().Inline(true).
			Reverse(true),
	}

	m.BufView = m.Buffer.View(0, 0)
	m.Cursor = m.BufView.Cursor(0, 0)

	return m
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
		m.WinStyle.Width(msg.Width - 2).Height(msg.Height - 2)
		m.BufView.Resize(msg.Height - 4)
	}

	return m, nil
}

func (m Model) onKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.Cancel):
		return m, tea.Quit
	}

	switch msg.Type {
	case tea.KeyRunes, tea.KeySpace:
		m.Cursor.Write([]byte(string(msg.Runes)))
	case tea.KeyEnter:
		m.Cursor.WriteByte('\n')
	case tea.KeyTab:
		m.Cursor.WriteByte('\t')
	case tea.KeyBackspace:
		m.Cursor.Delete(-1)
	}

	return m, nil
}

func (m Model) View() string {
	if (m.Size.Width == 0) || (m.Size.Height == 0) {
		return ""
	}

	var view strings.Builder

	{
		cursor := m.CursorStyle.Render(" ")
		win := m.WinStyle.Render(m.BufView.String() + cursor)
		view.WriteString(win)
	}

	return view.String()
}
