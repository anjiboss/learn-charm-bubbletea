package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
}

func (m *Model) Init() tea.Cmd {

	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m *Model) View() string {

	return ""
}

func main() {
	err := tea.NewProgram(&Model{}).Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
