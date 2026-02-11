package events

import (
	"github.com/charmbracelet/huh"
)

type FormSubmitted struct {
	Form *huh.Form
}

// Attaches "Event" interface to FormSubmitted (also implements tea.Msg inherintly)
func (s FormSubmitted) isEvent() {}
