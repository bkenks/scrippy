package scripts

import (
	efc "github.com/bkenks/scrippy/internal/scripts/editfishconfig"
	"github.com/charmbracelet/bubbles/list"
)

func All() []list.Item {
	return []list.Item{
		efc.New(),
	}
}
