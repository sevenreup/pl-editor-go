package textarea

import (
	"fmt"
	"sevenreup/pl-editor-go/src/components/utils"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
			case tea.KeyLeft:
				{
					m.Cursor.Col = utils.Clamp(m.Cursor.Col-1, 0, len(m.lines[m.Cursor.Line]))
				}
			case tea.KeyRight:
				{
					m.Cursor.Col = utils.Clamp(m.Cursor.Col+1, 0, len(m.lines[m.Cursor.Line]))
				}
			case tea.KeyBackspace:
				{
					m.deleteBeforeCursor()
				}
			case tea.KeyDelete:
				{
					m.deleteAfterCursor()
				}
			case tea.KeyEnter:
				{
					m.lines = append(m.lines, "")
					m.Cursor.Line = len(m.lines) - 1
					m.Cursor.Col = 0
				}
			default:
				{
					m.lines[m.Cursor.Line] = m.lines[m.Cursor.Line][:m.Cursor.Col] + string(msg.String()) + m.lines[m.Cursor.Line][m.Cursor.Col:]
					m.Cursor.Col++
				}
			}
		}
	}
	return m, nil
}

func (m *Model) deleteBeforeCursor() {
	m.Cursor.Col = utils.Clamp(m.Cursor.Col, 0, len(m.lines[m.Cursor.Line]))
	if m.Cursor.Col <= 0 {
		m.mergeLineAbove(m.Cursor.Line)
		return
	}
	if len(m.lines[m.Cursor.Line]) > 0 {
		m.lines[m.Cursor.Line] = m.lines[m.Cursor.Line][:len(m.lines[m.Cursor.Line])-1]
		m.Cursor.Col--
	}
}

func (m *Model) deleteAfterCursor() {}

func (m *Model) mergeLineAbove(line int) {
	if line <= 0 {
		return
	}
	m.lines[line-1] += m.lines[line]
	m.lines = append(m.lines[:line], m.lines[line+1:]...)

	m.Cursor.Col = len(m.lines[line-1])
	m.Cursor.Line = m.Cursor.Line - 1
}

func (m Model) View() string {
	var s strings.Builder
	var style lipgloss.Style

	for idx, line := range m.lines {
		s.WriteString(style.Width(3).Foreground(lipgloss.Color("#999999")).Render(fmt.Sprint(idx+1)) + " ")
		s.WriteString(line + "\n")
	}
	return fmt.Sprint(s.String())
}
