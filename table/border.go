package table

type Border struct {
	Top         string
	Left        string
	Right       string
	Bottom      string
	TopRight    string
	TopLeft     string
	BottomRight string
	BottomLeft  string

	TopJunction    string
	LeftJunction   string
	RightJunction  string
	BottomJunction string

	InnerJunction string

	InnerDivider string
}

var (
	borderDefault = Border{
		Top:    "━",
		Left:   "┃",
		Right:  "┃",
		Bottom: "━",

		TopRight:    "┓",
		TopLeft:     "┏",
		BottomRight: "┛",
		BottomLeft:  "┗",

		TopJunction:    "┳",
		LeftJunction:   "┣",
		RightJunction:  "┫",
		BottomJunction: "┻",
		InnerJunction:  "╋",

		InnerDivider: "┃",
	}

	noBorder = Border{
		Top:    " ",
		Left:   " ",
		Right:  " ",
		Bottom: " ",

		TopRight:    " ",
		TopLeft:     " ",
		BottomRight: " ",
		BottomLeft:  " ",

		TopJunction:    " ",
		LeftJunction:   " ",
		RightJunction:  " ",
		BottomJunction: " ",
		InnerJunction:  " ",

		InnerDivider: " ",
	}

	borderRounded = Border{
		Top:    "─",
		Left:   "│",
		Right:  "│",
		Bottom: "─",

		TopRight:    "╮",
		TopLeft:     "╭",
		BottomRight: "╯",
		BottomLeft:  "╰",

		TopJunction:    "┬",
		LeftJunction:   "├",
		RightJunction:  "┤",
		BottomJunction: "┴",
		InnerJunction:  "┼",

		InnerDivider: "│",
	}
)
