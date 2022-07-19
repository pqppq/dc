package keybinds

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/pqppq/dc/connection"
)

func getSchema(g *gocui.Gui, v *gocui.View) error {
	out, err := g.View("schema")
	if err != nil {
		return err
	}
	_, y := v.Cursor()
	line, _ := v.Line(y)

	name := strings.Split(line, " ")[1]
	schema, err := connection.GetSchema(name)
	if err != nil {
		return err
	}
	// v.Clear()
	for _, line = range schema {
		fmt.Fprintln(out, "▸", line)
	}

	return nil
}
