package events

type ScriptComplete struct {
	Err        error
	ScriptName string
}

// Attaches "Event" interface to StateChanged (also implements tea.Msg inherintly)
func (s ScriptComplete) isEvent() {}
