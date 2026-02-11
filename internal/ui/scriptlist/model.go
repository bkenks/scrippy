package scriptlist

import (
	"github.com/bkenks/scrippy/internal/commands"
	"github.com/bkenks/scrippy/internal/constants"
	"github.com/bkenks/scrippy/internal/domain"
	"github.com/bkenks/scrippy/internal/scripts"
	"github.com/bkenks/scrippy/internal/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Interface
//
// uiScriptList (tea.Model):
//	- Model (UI) for listing the repos from ghq and allowing the user to open them with Lazygit

type Model struct {
	ScriptList list.Model
}

func New() *Model {
	w, h := sizeBuffer()

	newList := list.New(
		scripts.All(),             // []list.Item containing the parsed list of repos from ghq
		list.NewDefaultDelegate(), // Default list.Item styling
		w, h)                      // Width & Height
	newList.Title = "Scrippy"
	newList.AdditionalShortHelpKeys = constants.ScriptListKeyMap.HelpBinds(constants.Short)
	newList.AdditionalFullHelpKeys = constants.ScriptListKeyMap.HelpBinds(constants.Full)

	return &Model{
		ScriptList: newList,
	}
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	/////////////////////////////////////
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		w, h := sizeBuffer()
		m.ScriptList.SetSize(w, h)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.ScriptListKeyMap.Run): // action: run script with new parameter selection
			selectedScript := m.ScriptList.SelectedItem() // type list.Item

			if scriptWithType, ok := selectedScript.(domain.Script); ok {
				scriptWithType.ResetForm() // reset the form to clear any previous state before passing it to the opts model

				cmds = append(cmds,
					commands.ScriptSelected(scriptWithType),
				)
			}
		case key.Matches(msg, constants.ScriptListKeyMap.RunPrevParams): // action: rerun script with past parameter selection
			selectedScript := m.ScriptList.SelectedItem() // type list.Item

			if scriptWithType, ok := selectedScript.(domain.Script); ok {
				// Don't reset the form to pass previous form to the opts model

				cmds = append(cmds,
					commands.ScriptSelected(scriptWithType),
				)
			}
		}
	}
	/////////////////////////////////////

	/////////////////////////////////////
	// Output
	m.ScriptList, cmd = m.ScriptList.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
	// End "Output"
	/////////////////////////////////////
}

func sizeBuffer() (width, height int) {
	x, y := styles.DocStyle.GetFrameSize()
	widthBuffer := constants.WindowSize.Width - x
	heightBuffer := constants.WindowSize.Height - y
	return widthBuffer, heightBuffer
}

func (m *Model) View() string { return m.ScriptList.View() }

// End "Interface"
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
