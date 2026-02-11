package opts

import (
	"github.com/bkenks/scrippy/internal/commands"
	"github.com/bkenks/scrippy/internal/constants"
	"github.com/bkenks/scrippy/internal/domain"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form      *huh.Form
	submitted bool
	script    domain.Script
}

func NewOpts(script domain.Script) *Model {
	return &Model{
		form:      script.Form(),
		submitted: false,
		script:    script,
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.OptsKeyMap.Exit):
			cmds = append(cmds,
				commands.SetState(domain.StateMain),
			)
		}
	}

	if !m.submitted && m.form.State == huh.StateCompleted {
		m.submitted = true
		cmds = append(cmds,
			commands.SetState(domain.StateMain),
			commands.FormSubmitted(m.form),
		)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.form.View()
}
