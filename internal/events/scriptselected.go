package events

import "github.com/bkenks/scrippy/internal/domain"

type ScriptSelected struct {
	Script domain.Script
}

// Attaches "Event" interface to StateChanged (also implements tea.Msg inherintly)
func (s ScriptSelected) isEvent() {}
