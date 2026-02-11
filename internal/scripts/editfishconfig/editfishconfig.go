// Change this to match your struct name and make sure to call .New() in registry.go
// ex:
// package "yourscriptname"
// yourscriptname.New() -> internal/scripts/registry.go
package efc

import (
	"os/exec"

	"github.com/bkenks/scrippy/internal/events"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var title = "Edit Fish Config"
var desc = "Edit the configuration for your fish shell"

// Create a form for optional configuration of the script
func newForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("editor").
				Title("Choose your editor").
				Options(huh.NewOptions("nvim", "vscode")...),
		),
	)
}

// Place your scripting logic here. This will be executed when the script is run.
// You can use fields from the form with e.Form().Get("key") to access the values submitted by the user.
func (e *EFC) Exec() tea.Cmd {
	var editor string
	switch e.Form().GetString("editor") {
	case "nvim":
		editor = "nvim"
	case "vscode":
		editor = "code"
	default:
		editor = "nvim" // default to nvim if something goes wrong
	}

	cmdBuilder := exec.Command(editor, "/home/brian-kenkel/.config/fish/config.fish")

	cmd := tea.ExecProcess(cmdBuilder, execCallback)

	return cmd
}

// Name this whatever you want, but it should be descriptive of the script's purpose
// This represents the script itself
type EFC struct {
	title, desc string
	form        *huh.Form
}

///////////////////////////////////////////////////////////////////////////////////////
// Initializer
// Implements: domain.Script & list.Item

func New() *EFC {
	return &EFC{
		title: title,
		desc:  desc,
		form:  newForm(),
	}
}

func (e *EFC) Title() string       { return e.title }
func (e *EFC) Description() string { return e.desc }
func (e *EFC) FilterValue() string { return e.title }
func (e *EFC) Form() *huh.Form     { return e.form }
func (e *EFC) ResetForm() {
	e.form = newForm()
}

var execCallback tea.ExecCallback = func(err error) tea.Msg {
	return events.ScriptComplete{
		ScriptName: title,
		Err:        err,
	}
}

// End "Initializer"
///////////////////////////////////////////////////////////////////////////////////////
