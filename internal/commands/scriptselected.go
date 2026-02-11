package commands

import (
	"github.com/bkenks/scrippy/internal/domain"
	"github.com/bkenks/scrippy/internal/events"
	tea "github.com/charmbracelet/bubbletea"
)

func ScriptSelected(script domain.Script) tea.Cmd {
	return func() tea.Msg {
		return events.ScriptSelected{
			Script: script,
		}
	}
}
