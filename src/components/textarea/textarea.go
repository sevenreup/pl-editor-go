package textarea

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type cursor struct {
	Line int
	Col  int
}

type Model struct {
	lines  []string
	Cursor cursor
}

func New() *Model {
	return initModel()
}

func initModel() *Model {
	cursor := &cursor{
		Line: 0,
		Col:  0,
	}
	return &Model{
		lines:  []string{""},
		Cursor: *cursor,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.Type {
			case tea.KeyUp:
				{
					if m.Cursor.Line > 0 {
						m.Cursor.Line--
					}
				}

			case tea.KeyDown:
				{
					if m.Cursor.Line < len(m.lines)-1 {
						m.Cursor.Line++
					}
				}
			case tea.KeyBackspace:
				{
					line := m.lines[m.Cursor.Line]
					if len(line) > 0 {
						m.lines[m.Cursor.Line] = line[:len(line)-1]
					}
				}
			case tea.KeyEnter:
				{
					m.lines = append(m.lines, "")
					m.Cursor.Line = len(m.lines) - 1
				}
			default:
				{
					m.lines[m.Cursor.Line] += msg.String()
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := ""
	for _, line := range m.lines {
		s += line + "\n"
	}
	return fmt.Sprint(s)
}
