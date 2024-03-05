package styles

import "github.com/charmbracelet/lipgloss"

const (
	Purple    = lipgloss.Color("#BD93F9")
	Gray      = lipgloss.Color("#8B8B8B")
	LightGray = lipgloss.Color("#DADADA")
	Cyan      = lipgloss.Color("#7DCFFF")
	Green     = lipgloss.Color("#28D33F")
	AliceBlue = lipgloss.Color("#f0f8ff")
	Error     = lipgloss.Color("#FF5555")
)

var TimerStyle = lipgloss.NewStyle().Foreground(Gray).Align(lipgloss.Center).Italic(true)
var ErrorStyle = lipgloss.NewStyle().Foreground(Error).Align(lipgloss.Center).Foreground(Error).Bold(true)

func CreateHeaderCell(text string, color lipgloss.TerminalColor) string {
	return lipgloss.NewStyle().
		Foreground(color).
		Bold(true).
		Width(10).
		Align(lipgloss.Center).
		Render(text)
}
