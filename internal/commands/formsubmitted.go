package commands

import (
	"github.com/bkenks/scrippy/internal/events"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func FormSubmitted(form *huh.Form) tea.Cmd {
	return func() tea.Msg {
		return events.FormSubmitted{
			Form: form,
		}
	}
}
