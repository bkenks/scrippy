package commands

import (
	"github.com/bkenks/scrippy/internal/domain"
	"github.com/bkenks/scrippy/internal/events"
	tea "github.com/charmbracelet/bubbletea"
)

func SetState(state domain.SessionState) tea.Cmd {
	return func() tea.Msg {
		return events.StateChanged{
			State: state,
		}
	}
	
}