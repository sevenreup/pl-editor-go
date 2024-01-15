package main

import (
	"fmt"
	"os"
	"sevenreup/pl-editor-go/src/components/textarea"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textarea textarea.Model
}

func main() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initModel() *model {
	area := textarea.New()

	return &model{
		textarea: *area,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.Type {
			case tea.KeyCtrlC:
				{
					return m, tea.Quit
				}
			default:
				{

				}
			}
		}
	}
	md, cm := m.textarea.Update(msg)
	// using type assertion to get the textarea.Model
	m.textarea = md.(textarea.Model)
	cmd = cm
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprint(m.textarea.View())
}
