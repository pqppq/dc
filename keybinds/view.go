package keybinds

import (
	"github.com/jroimartin/gocui"
	"github.com/pqppq/dc/layout"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// move to next view
func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (layout.CURRENT_VIEW + 1) % len(layout.VIEW_NAMES)
	name := layout.VIEW_NAMES[nextIndex]

	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	layout.CURRENT_VIEW = nextIndex
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
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

func cursorUp(g *gocui.Gui, v *gocui.View) error {
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
