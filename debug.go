package gui

import (
	"fmt"
	"reflect"
	"strings"
)

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
	case *Container:
		wType = "C"
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
	case *Button:
		extra = "text:\"" + w.text + "\""
	}

	return fmt.Sprintf("(%s)%s %s %s %s", wType, wId, widget.Bounds().ToString(), flags, extra)
}

func ContainerToTreeDebugString(c IContainer) string {
	var depth string
	return toString(c, depth)
}

func toString(node IWidget, depth string) string {
	s := WidgetToDebugString(node) + "\n"
	container, ok := node.(*Container)
	if ok {
		for index, _ := range container.Children() {
			child := container.children[index]
			s += fmt.Sprintf("%s `--", depth)
			c := " "
			if index < container.ChildrenCount()-1 {
				c = "|"
			}
			depth += " " + c + "  "
			s += toString(child, depth)
			depth = depth[:len(depth)-4]
		}
	}
	return s
}