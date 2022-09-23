// Package anvil implements an Acme-style layout system.
package anvil

type Layout struct {
	cols []Column
}

func (layout Layout) View() string {
	panic("Not implemented.")
}

type Column struct {
	wins []Window
}

type Window struct {
	cmds string
	data string
}
