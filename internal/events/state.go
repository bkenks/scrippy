package events

import (
	"github.com/bkenks/scrippy/internal/domain"
)

type StateChanged struct{
	State domain.SessionState
}

// Attaches "Event" interface to StateChanged (also implements tea.Msg inherintly)
func (s StateChanged) isEvent() {}