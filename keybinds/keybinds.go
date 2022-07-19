package keybinds

import (
	"log"

	"github.com/jroimartin/gocui"
)

type mapping struct {
	viewname string
	key      interface{}
	mod      gocui.Modifier
	handler  func(*gocui.Gui, *gocui.View) error
}

var mappings = []mapping{
	{
		"", gocui.KeyCtrlC, gocui.ModNone, quit,
	},
	{
		"", gocui.KeyTab, gocui.ModNone, nextView,
	},
	{
		"", 'j', gocui.ModNone, cursorDown,
	},
	{
		"", 'k', gocui.ModNone, cursorUp,
	},
	{
		"", gocui.KeyEnter, gocui.ModNone, getSchema,
	},
	// TODO
	// Resize Window Action
}

func SetKeybinds(g *gocui.Gui) error {

	for _, mapping := range mappings {
		if err := g.SetKeybinding(mapping.viewname, mapping.key, mapping.mod, mapping.handler); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
