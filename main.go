package main

import (
	"log"

	"github.com/pqppq/dc/keybinds"
	"github.com/pqppq/dc/layout"

	"github.com/jroimartin/gocui"
)

func main() {
	// create gui
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.Cursor = true
	g.SelFgColor = gocui.ColorWhite

	// set layout
	layout.SetLayout(g)
	// set keybindinds
	keybinds.SetKeybinds(g)

	// start main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}

}
