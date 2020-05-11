package gui

import (
	"fmt"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"reflect"
	"strings"
)

func DimensionToString(dimension utils.Dimension) string {
	switch dimension.Unit {
	case utils.PixelUnit:
		return fmt.Sprintf("%dpx", dimension.Value)
	case utils.PercentageUnit:
		return fmt.Sprintf("%d%%", dimension.Value)
	}

	return "-"
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

	// Bounds
	boundsString := fmt.Sprintf(
		"%s[%s,%s]",
		widget.Bounds().ToString(),
		DimensionToString(widget.DimensionH()),
		DimensionToString(widget.DimensionV()),
	)

	return fmt.Sprintf("(%s)%s %s %s %s", wType, wId,  boundsString, flags, extra)
}

func ContainerToTreeDebugString(c IContainer) string {
	return containerToTreeDebugString(c, "")
}

func containerToTreeDebugString(node IWidget, depth string) string {
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
			s += containerToTreeDebugString(child, depth)
			depth = depth[:len(depth)-4]
		}
	}
	bc, ok := node.(*BoxContainer)
	if ok {
		for index, _ := range bc.Children() {
			child := bc.children[index]
			s += fmt.Sprintf("%s `--", depth)
			c := " "
			if index < bc.ChildrenCount()-1 {
				c = "|"
			}
			depth += " " + c + "  "
			s += containerToTreeDebugString(child, depth)
			depth = depth[:len(depth)-4]
		}
	}
	return s
}
