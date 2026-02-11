package styles

import "github.com/charmbracelet/lipgloss"

const maxWidth = 80

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

var (
	DarkPink   = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	DullGrey   = lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}
	Purple     = lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}
	DarkPurple = lipgloss.Color("62")
	White      = lipgloss.Color("230")

	ButtonStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Margin(0, 1)

	SelectedButton = ButtonStyle.
			Background(DarkPurple).
			Foreground(White).
			Bold(true)

	UnselectedButton = ButtonStyle.
				Background(DullGrey).
				Foreground(lipgloss.Color("250"))

	DocStyle = lipgloss.NewStyle().
			Margin(3, 1)

	DialogStyle = lipgloss.NewStyle().
			Padding(0, 6, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(DullGrey)

	DialogTitleStyle = lipgloss.NewStyle().Bold(true)

	TitleStyle = lipgloss.NewStyle().
			Background(DarkPurple).
			Foreground(White).
			Padding(0, 1).
			Bold(true).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			MarginBottom(1)
)
