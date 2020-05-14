package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"testing"
)

type WidgetTest struct {
	gui.Widget
}

func TestWidget(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
	w := &WidgetTest{}
	g.Screen().AddChild(w)

	// Minimum / Maximum
	w.SetMinimumWidth(100)
	w.SetMaximumWidth(300)
	w.SetWidth(120)
	assertIntEqual(t, w.Bounds().W, 120)
	w.SetWidth(80)
	assertIntEqual(t, w.Bounds().W, 100)
	w.SetWidth(400)
	assertIntEqual(t, w.Bounds().W, 300)

	w.SetMinimumHeight(400)
	w.SetMaximumHeight(800)
	w.SetHeight(500)
	assertIntEqual(t, w.Bounds().H, 500)
	w.SetHeight(380)
	assertIntEqual(t, w.Bounds().H, 400)
	w.SetHeight(900)
	assertIntEqual(t, w.Bounds().H, 800)
}
