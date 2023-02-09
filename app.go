package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := menu{
		options: []menuItem{
			menuItem{
				text:    "new check-in",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			menuItem{
				text:    "view check-ins",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
		},
	}

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
