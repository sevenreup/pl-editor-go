package info

import (
	"fmt"
	"sevenreup/pl-editor-go/src/components/textarea"
)

func DrawBottomInfo(m textarea.Model) string {
	return fmt.Sprintf("Line: %d, Col: %d", m.Cursor.Line+1, m.Cursor.Col+1)
}
