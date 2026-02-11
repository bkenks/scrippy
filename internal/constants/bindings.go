package constants

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////

type HelpType int

const (
	Short HelpType = iota
	Full
)

var unsetText = "not set"

type keyMap interface {
	HelpBinds()
}

func SetOnHelpType(helpType HelpType, bind key.Binding, shortHelp string, fullHelp string) key.Binding {
	bindWithHelp := bind

	switch helpType {
	case Short:
		bindWithHelp.SetHelp(bind.Help().Key, shortHelp)
	case Full:
		bindWithHelp.SetHelp(bind.Help().Key, fullHelp)
	}
	return bindWithHelp
}

//// End "Helpers"
///////////////////////////////////////////////////////

///////////////////////////////////////////////////////
//// Default Key Map

type defaultKeyMap struct {
	Select key.Binding
	Exit   key.Binding
}

var DefaultKeyMap = defaultKeyMap{
	Select: key.NewBinding(
		key.WithKeys(
			tea.KeyEnter.String(),
			tea.KeySpace.String(),
		),
		key.WithHelp(
			tea.KeyEnter.String()+"/"+tea.KeySpace.String(),
			unsetText,
		),
	),
	Exit: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp(
			tea.KeyEsc.String(),
			unsetText,
		),
	),
}

func (k defaultKeyMap) HelpBinds(helpType HelpType) func() []key.Binding {
	bindsWithHelp := []key.Binding{
		SetOnHelpType(
			helpType,             // Short or Full Help
			DefaultKeyMap.Select, // key.Binding
			"select",             // Short Help
			"select",             // Full Help
		),
		SetOnHelpType(
			helpType,
			DefaultKeyMap.Exit,
			"exit",
			"exit",
		),
	}

	return func() []key.Binding { return bindsWithHelp }
}

//// End "Default Key Map"
///////////////////////////////////////////////////////

///////////////////////////////////////////////////////
//// Script List Key Map

type scriptListKeyMap struct {
	Run           key.Binding
	RunPrevParams key.Binding
}

var ScriptListKeyMap = scriptListKeyMap{
	Run: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", unsetText),
	),
	RunPrevParams: key.NewBinding(
		key.WithKeys(tea.KeyCtrlR.String()),
		key.WithHelp("ctrl+r", unsetText),
	),
}

func (k scriptListKeyMap) HelpBinds(helpType HelpType) func() []key.Binding {
	bindsWithHelp := []key.Binding{
		SetOnHelpType(
			helpType,              // Short or Full Help
			ScriptListKeyMap.Run,  // key.Binding
			"run",                 // Short Help
			"run with new params", // Full Help
		),
		SetOnHelpType(
			helpType,
			ScriptListKeyMap.RunPrevParams,
			"run prev params",
			"run with previous params",
		),
	}

	return func() []key.Binding { return bindsWithHelp }
}

//// End "Script List Key Map"
///////////////////////////////////////////////////////

///////////////////////////////////////////////////////
//// Opts Key Map

type optsKeyMap struct {
	Exit key.Binding
}

var OptsKeyMap = optsKeyMap{
	Exit: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp(tea.KeyEsc.String(), unsetText),
	),
}

func (k optsKeyMap) HelpBinds(helpType HelpType) func() []key.Binding {
	bindsWithHelp := []key.Binding{
		SetOnHelpType(
			helpType,              // Short or Full Help
			OptsKeyMap.Exit,       // key.Binding
			"back",                // Short Help
			"back to script list", // Full Help
		),
	}

	return func() []key.Binding { return bindsWithHelp }
}

//// Opts Key Map
///////////////////////////////////////////////////////
