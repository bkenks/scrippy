package app

import (
	"github.com/bkenks/scrippy/internal/commands"
	"github.com/bkenks/scrippy/internal/constants"
	"github.com/bkenks/scrippy/internal/domain"
	"github.com/bkenks/scrippy/internal/events"
	"github.com/bkenks/scrippy/internal/styles"
	"github.com/bkenks/scrippy/internal/ui/opts"
	"github.com/bkenks/scrippy/internal/ui/scriptlist"
	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Interface: tea.Model
//
// ModelManager:
//	- Model for managing sub-Models (i.e other UI/Views/Screens)

type ModelManager struct {
	state domain.SessionState

	scripts  []domain.Script
	selected domain.Script

	main   scriptlist.Model
	opts   opts.Model
	active tea.Model
}

func New() *ModelManager {
	m := &ModelManager{
		main: *scriptlist.New(), // Main Model (List)
	}

	m.active = &m.main
	return m
}

func (m *ModelManager) Init() tea.Cmd { return nil }

func (m *ModelManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	/////////////////////////////////////
	// UI Manager
	switch msg := msg.(type) {

	//// Reactive Window Sizing
	case tea.WindowSizeMsg:
		constants.WindowSize = msg

	case events.Event:
		switch msg := msg.(type) {
		//// State Manager
		case events.StateChanged:
			m.state = msg.State

			// Initialization for each state
			switch m.state {
			case domain.StateMain:
				m.active = &m.main
			case domain.StateOpts:
				m.active = &m.opts
			}

		case events.ScriptSelected:
			m.selected = msg.Script

			m.opts = *opts.NewOpts(m.selected)
			// Handle script selection (e.g., transition to script detail view)
			cmds = append(cmds,
				commands.SetState(domain.StateOpts),
			)
		case events.FormSubmitted:
			cmds = append(cmds,
				m.selected.Exec(),
			)
		}
	}
	// End "UI Manager"
	/////////////////////////////////////

	var cmd tea.Cmd
	m.active, cmd = m.active.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *ModelManager) View() string {
	return styles.DocStyle.Render(m.active.View())
}

// End "Interface: tea.Model"
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
