package gui

import (
	"fmt"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"reflect"
	"strings"
)

func ColorToString(f utils.Color) string {
	return fmt.Sprintf("{%.2f,%.2f,%.2f,%.2f}", f[0], f[1], f[2], f[3])
}

func InsetsToString(i utils.Insets) string {
	return fmt.Sprintf("{top:%d,right:%d,bottom:%d,left:%d}", i.Top, i.Right, i.Bottom, i.Left)
}

func WidgetToDebugString(str interface{}) string {
	widget := str.(IWidget)
	// Widget id
	wId := ""
	if widget.Id() != "" {
		wId = " '" + widget.Id() + "'"
	}
	// Struct type
	wType := ""
	switch str.(type) {
	case *Screen:
		wType = "S"
	case *Container:
		wType = "C"
	case *BoxContainer:
		wType = "BC"
	case *Label:
		wType = "L"
	case *Button:
		wType = "B"
	default:
		wType = strings.Replace(reflect.TypeOf(widget).String(), "*gui.", "", -1)
	}

	// Flags
	flags := ""
	if widget.Enabled() {
		flags += "E"
	}
	if widget.Visible() {
		flags += "V"
	}

	// Extra content
	extra := ""
	switch w := str.(type) {
	case *Label:
		extra = "text:\"" + w.text + "\""
	case *Button:
		extra = "text:\"" + w.text + "\""
	}

	return fmt.Sprintf("(%s)%s %s %s %s", wType, wId, widget.Bounds().ToString(), flags, extra)
}

func ContainerToTreeDebugString(c IContainer) string {
	return containerToTreeDebugString(c, "")
}

func containerToTreeDebugString(node IWidget, depth string) string {
	s := WidgetToDebugString(node) + "\n"

	container, ok := node.(IContainer)
	if ok {
		for index, _ := range container.Children() {
			child := container.Children()[index]
			s += fmt.Sprintf("%s `--", depth)
			c := " "
			if index < container.ChildrenCount()-1 {
				c = "|"
			}
			depth += " " + c + "  "
			s += containerToTreeDebugString(child, depth)
			depth = depth[:len(depth)-4]
		}
	}
	return s
}
