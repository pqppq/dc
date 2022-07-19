package layout

import (
	"fmt"
	"strings"

	"github.com/pqppq/dc/connection"

	"github.com/jroimartin/gocui"
)

// view dimension
type dimension struct {
	x0, y0, x1, y1 int // left top right bottom
}

var (
	CURRENT_VIEW = 0
	VIEW_NAMES   = []string{"connection", "query", "schema", "result"}
)

// return dimension of view
func getDimension(name string, maxX, maxY int) (dimension, error) {
	x, y := int(maxX*2/10), int(maxY*2/10)

	if name == "connection" {
		return dimension{
			0, 0, x, y,
		}, nil
	}
	if name == "query" {
		return dimension{
			x + 1, 0, maxX - 1, y,
		}, nil
	}
	if name == "schema" {
		return dimension{
			0, y + 1, x, maxY - 1,
		}, nil
	}
	if name == "result" {
		return dimension{
			x + 1, y + 1, maxX - 1, maxY - 1,
		}, nil
	}
	return dimension{}, gocui.ErrUnknownView
}

// configure view attributes
func configureView(v *gocui.View) {
	name := v.Name()
	if name == "connection" {
		v.Title = " connection "
		v.Wrap = false
		v.Editable = false
		v.Highlight = true
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorWhite
		for _, db := range connection.GetDBs() {
			width, _ := v.Size()
			fmt.Fprintln(v, "* "+db.Name+strings.Repeat(" ", width)) // padding with white space
		}
	}
	if name == "query" {
		v.Title = " query "
		v.Wrap = false
		v.Editable = false
	}
	if name == "schema" {
		v.Title = " schema "
		v.Wrap = false
		v.Editable = false
	}
	if name == "result" {
		v.Title = " result "
		v.Wrap = false
		v.Editable = false
	}
	return
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// DELETE
	// if out, err := g.View("result"); err == nil {
	// 	fmt.Fprintln(out, maxX, maxY)
	// 	a, _ := getDimension("connection", maxX, maxY)
	// 	fmt.Fprintln(out, a)
	// }

	for _, name := range VIEW_NAMES {
		dim, err := getDimension(name, maxX, maxY)
		if err != nil {
			return err
		}
		v, err := g.SetView(name, dim.x0, dim.y0, dim.x1, dim.y1)
		if err != gocui.ErrUnknownView {
			return err
		}
		configureView(v)
		if _, err := g.SetCurrentView("connection"); err != nil {
			return err
		}
	}
	if name := g.CurrentView().Name(); name == "query" || name == "result" {
		g.Cursor = true
	} else {
		g.Cursor = false
	}
	return nil
}

func SetLayout(g *gocui.Gui) {
	g.SetManagerFunc(layout)
}

// move to next view
func NextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (CURRENT_VIEW + 1) % len(VIEW_NAMES)
	name := VIEW_NAMES[nextIndex]

	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	CURRENT_VIEW = nextIndex
	return nil
}

func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		l := len(v.BufferLines()) - 1
		ln := cy - oy
		_, ly := v.Size()

		// case when buffer lines is less than view height
		if l < ly {
			if ln+1 < l {
				if err := v.SetCursor(cx, cy+1); err != nil {
					return err
				}
			}
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			if ly+oy < l {
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
