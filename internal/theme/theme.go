package theme

import "github.com/charmbracelet/lipgloss"

var (
	red      = lipgloss.Color("#ff8272")
	green    = lipgloss.Color("#b4fa72")
	yellow   = lipgloss.Color("#fefdc2")
	blue     = lipgloss.Color("#a5d5fe")
	neonBlue = lipgloss.Color("#1ec9ff")
	pink     = lipgloss.Color("#ff8ffd")
	magenta  = lipgloss.Color("#d0d1fe")
	white    = lipgloss.Color("#f1f1f1")
	darkGray = lipgloss.Color("#616161")
)

var (
	Title              = lipgloss.NewStyle().Foreground(white).Bold(true)
	Foreground         = lipgloss.NewStyle().Foreground(white)
	AccentRed          = lipgloss.NewStyle().Bold(true).Foreground(red)
	AccentBlue         = lipgloss.NewStyle().Bold(true).Foreground(neonBlue)
	AccentGreen        = lipgloss.NewStyle().Bold(true).Foreground(green)
	RoundSelector      = lipgloss.NewStyle().Foreground(red).Render("●")
	Selector           = lipgloss.NewStyle().Foreground(red).Render("❯")
	Dimmed             = lipgloss.NewStyle().Faint(true)
	Highlight          = lipgloss.NewStyle().Foreground(green)
	Accent             = lipgloss.NewStyle().Foreground(pink)
	SelectedOption     = lipgloss.NewStyle().Bold(true).Foreground(green)
	NonSelectedOption  = lipgloss.NewStyle().Faint(true)
	Warning            = lipgloss.NewStyle().Bold(true).Foreground(yellow)
	TableCellContent   = lipgloss.NewStyle().Foreground(white)
	TableColumnContent = lipgloss.NewStyle().Foreground(white).Bold(true).AlignHorizontal(lipgloss.Center).Padding(0, 1, 0, 1)
)
