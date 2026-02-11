package domain

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Script interface {
	Title() string
	Description() string
	FilterValue() string
	Form() *huh.Form
	ResetForm()
	Exec() tea.Cmd
}
