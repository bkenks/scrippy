package main

import (
	"fmt"
	"os"

	"github.com/bkenks/scrippy/internal/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	scrippy := app.New()
	p := tea.NewProgram(scrippy, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
