package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type cursor struct {
	line int
	col  int
}
type model struct {
	lines  []string
	cursor cursor
}

func main() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initModel() *model {
	cursor := &cursor{
		line: 0,
		col:  0,
	}
	return &model{
		lines:  []string{""},
		cursor: *cursor,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.Type {
			case tea.KeyCtrlC:
				{
					return m, tea.Quit
				}
			case tea.KeyUp:
				{
					if m.cursor.line > 0 {
						m.cursor.line--
					}
				}

			case tea.KeyDown:
				{
					if m.cursor.line < len(m.lines)-1 {
						m.cursor.line++
					}
				}
			case tea.KeyBackspace:
				{
					line := m.lines[m.cursor.line]
					if len(line) > 0 {
						m.lines[m.cursor.line] = line[:len(line)-1]
					}
				}
			case tea.KeyEnter:
				{
					m.lines = append(m.lines, "")
					m.cursor.line = len(m.lines) - 1
				}
			default:
				{
					m.lines[m.cursor.line] += msg.String()
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""
	for _, line := range m.lines {
		s += line + "\n"
	}
	return fmt.Sprint(s)
}
