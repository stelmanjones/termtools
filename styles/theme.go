package styles

import (
	"github.com/muesli/termenv"
)

var s termenv.Style

var (
	red      = termenv.RGBColor("#ff8272")
	green    = termenv.RGBColor("#b4fa72")
	yellow   = termenv.RGBColor("#fefdc2")
	blue     = termenv.RGBColor("#a5d5fe")
	neonBlue = termenv.RGBColor("#1ec9ff")
	pink     = termenv.RGBColor("#ff8ffd")
	magenta  = termenv.RGBColor("#d0d1fe")
	white    = termenv.RGBColor("#f1f1f1")
	darkGray = termenv.RGBColor("#616161")
)

var (
	Title             = s.Foreground(white).Bold()
	Foreground        = s.Foreground(white)
	AccentRed         = s.Bold().Foreground(red)
	AccentBlue        = s.Bold().Foreground(neonBlue)
	AccentGreen       = s.Bold().Foreground(green)
	Selector          = s.Foreground(red).Styled("❯")
	Dimmed            = s.Faint()
	Highlight         = s.Foreground(green)
	Accent            = s.Foreground(pink)
	SelectedOption    = s.Bold().Foreground(green)
	NonSelectedOption = s.Faint()
	Warning           = s.Foreground(yellow)
)

/*
// Theme is a collection of styles for components of the form.
// Themes can be applied to a form using the WithTheme option.
type Theme struct {
	Form           lipgloss.Style
	Group          lipgloss.Style
	FieldSeparator lipgloss.Style
	Blurred        FieldStyles
	Focused        FieldStyles
	Help           help.Styles
}

// copy returns a copy of a theme with all children styles copied.
func (t Theme) copy() Theme {
	return Theme{
		Form:           t.Form.Copy(),
		Group:          t.Group.Copy(),
		FieldSeparator: t.FieldSeparator.Copy(),
		Blurred:        t.Blurred.copy(),
		Focused:        t.Focused.copy(),
		Help: help.Styles{
			Ellipsis:       t.Help.Ellipsis.Copy(),
			ShortKey:       t.Help.ShortKey.Copy(),
			ShortDesc:      t.Help.ShortDesc.Copy(),
			ShortSeparator: t.Help.ShortSeparator.Copy(),
			FullKey:        t.Help.FullKey.Copy(),
			FullDesc:       t.Help.FullDesc.Copy(),
			FullSeparator:  t.Help.FullSeparator.Copy(),
		},
	}
}

// FieldStyles are the styles for input fields.
type FieldStyles struct {
	Base           lipgloss.Style
	Title          lipgloss.Style
	Description    lipgloss.Style
	ErrorIndicator lipgloss.Style
	ErrorMessage   lipgloss.Style

	// Select styles.
	SelectSelector lipgloss.Style // Selection indicator
	Option         lipgloss.Style // Select options
	NextIndicator  lipgloss.Style
	PrevIndicator  lipgloss.Style

	// FilePicker styles.
	Directory lipgloss.Style
	File      lipgloss.Style

	// Multi-select styles.
	MultiSelectSelector lipgloss.Style
	SelectedOption      lipgloss.Style
	SelectedPrefix      lipgloss.Style
	UnselectedOption    lipgloss.Style
	UnselectedPrefix    lipgloss.Style

	// Textinput and teatarea styles.
	TextInput TextInputStyles

	// Confirm styles.
	FocusedButton lipgloss.Style
	BlurredButton lipgloss.Style

	// Card styles.
	Card      lipgloss.Style
	NoteTitle lipgloss.Style
	Next      lipgloss.Style
}

// TextInputStyles are the styles for text inputs.
type TextInputStyles struct {
	Cursor      lipgloss.Style
	Placeholder lipgloss.Style
	Prompt      lipgloss.Style
	Text        lipgloss.Style
}

// copy returns a copy of a TextInputStyles with all children styles copied.
func (t TextInputStyles) copy() TextInputStyles {
	return TextInputStyles{
		Cursor:      t.Cursor.Copy(),
		Placeholder: t.Placeholder.Copy(),
		Prompt:      t.Prompt.Copy(),
		Text:        t.Text.Copy(),
	}
}

// copy returns a copy of a FieldStyles with all children styles copied.
func (f FieldStyles) copy() FieldStyles {
	return FieldStyles{
		Base:                f.Base.Copy(),
		Title:               f.Title.Copy(),
		Description:         f.Description.Copy(),
		ErrorIndicator:      f.ErrorIndicator.Copy(),
		ErrorMessage:        f.ErrorMessage.Copy(),
		SelectSelector:      f.SelectSelector.Copy(),
		NextIndicator:       f.NextIndicator.Copy(),
		PrevIndicator:       f.PrevIndicator.Copy(),
		Option:              f.Option.Copy(),
		Directory:           f.Directory.Copy(),
		File:                f.File.Copy(),
		MultiSelectSelector: f.MultiSelectSelector.Copy(),
		SelectedOption:      f.SelectedOption.Copy(),
		SelectedPrefix:      f.SelectedPrefix.Copy(),
		UnselectedOption:    f.UnselectedOption.Copy(),
		UnselectedPrefix:    f.UnselectedPrefix.Copy(),
		FocusedButton:       f.FocusedButton.Copy(),
		BlurredButton:       f.BlurredButton.Copy(),
		TextInput:           f.TextInput.copy(),
		Card:                f.Card.Copy(),
		NoteTitle:           f.NoteTitle.Copy(),
		Next:                f.Next.Copy(),
	}
}

const (
	buttonPaddingHorizontal = 2
	buttonPaddingVertical   = 0
)

// ThemeBase returns a new base theme with general styles to be inherited by
// other themes.
func ThemeBase() *Theme {
	var t Theme

	t.FieldSeparator = lipgloss.NewStyle().SetString("\n\n")

	button := lipgloss.NewStyle().
		Padding(buttonPaddingVertical, buttonPaddingHorizontal).
		MarginRight(1)

	// Focused styles.
	f := &t.Focused
	f.Base = lipgloss.NewStyle().
		PaddingLeft(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true)
	f.Card = lipgloss.NewStyle().
		PaddingLeft(1)
	f.ErrorIndicator = lipgloss.NewStyle().
		SetString(" *")
	f.ErrorMessage = lipgloss.NewStyle().
		SetString(" *")
	f.SelectSelector = lipgloss.NewStyle().
		SetString("> ")
	f.NextIndicator = lipgloss.NewStyle().
		MarginLeft(1).
		SetString("→")
	f.PrevIndicator = lipgloss.NewStyle().
		MarginRight(1).
		SetString("←")
	f.MultiSelectSelector = lipgloss.NewStyle().
		SetString("> ")
	f.SelectedPrefix = lipgloss.NewStyle().
		SetString("[•] ")
	f.UnselectedPrefix = lipgloss.NewStyle().
		SetString("[ ] ")
	f.FocusedButton = button.Copy().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("7"))
	f.BlurredButton = button.Copy().
		Foreground(lipgloss.Color("7")).
		Background(lipgloss.Color("0"))
	f.TextInput.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	t.Help = help.New().Styles

	// Blurred styles.
	t.Blurred = f.copy()
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.MultiSelectSelector = lipgloss.NewStyle().SetString("  ")
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return &t
}

// ThemeCharm returns a new theme based on the Charm color scheme.
func ThemeCharm() *Theme {
	t := ThemeBase().copy()

	var (
		normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
		cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
		fuchsia  = lipgloss.Color("#F780E2")
		green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
		red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	)

	f := &t.Focused
	f.Base = f.Base.BorderForeground(lipgloss.Color("238"))
	f.Title.Foreground(indigo).Bold(true)
	f.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
	f.Directory.Foreground(indigo)
	f.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	f.ErrorIndicator.Foreground(red)
	f.ErrorMessage.Foreground(red)
	f.SelectSelector.Foreground(fuchsia)
	f.NextIndicator.Foreground(fuchsia)
	f.PrevIndicator.Foreground(fuchsia)
	f.Option.Foreground(normalFg)
	f.MultiSelectSelector.Foreground(fuchsia)
	f.SelectedOption.Foreground(green)
	f.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	f.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	f.UnselectedOption.Foreground(normalFg)
	f.FocusedButton.Foreground(cream).Background(fuchsia)
	f.Next = f.FocusedButton.Copy()
	f.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	f.TextInput.Cursor.Foreground(green)
	f.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	f.TextInput.Prompt.Foreground(fuchsia)

	t.Blurred = f.copy()
	t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return &t
}

*/
