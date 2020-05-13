package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type ButtonType int

const (
	ButtonTypeNormal ButtonType = iota
	ButtonTypeToggle
)

type Button struct {
	Label
	buttonType  ButtonType
	buttonGroup []*Button
	pressed     bool
}

func NewButton(text string) *Button {
	b := &Button{}
	widgetInit(b)
	b.buttonType = ButtonTypeNormal
	b.fontSize = 25
	b.padding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2}
	b.contentAlignmentH = utils.AlignmentHCenter
	b.contentAlignmentV = utils.AlignmentVCenter
	b.text = text
	return b
}

func NewToggleButton(text string) *Button {
	b := NewButton(text)
	b.buttonType = ButtonTypeToggle
	return b
}

// Getters
func (b *Button) Pressed() bool { return b.pressed }

func (b *Button) SetTheme(theme *Theme) {
	b.theme = theme
	b.font = theme.ButtonFont
	b.textColor = theme.ButtonTextColor
	b.fontSize = theme.ButtonFontSize
	b.padding = theme.ButtonPadding
	b.Measure()
}

func (b *Button) SetButtonGroup(buttonGroup []*Button) {
	b.buttonGroup = buttonGroup
}

func (b *Button) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if event == EventActionPress {
		if b.buttonType == ButtonTypeNormal {
			b.pressed = true
		} else {
			if b.buttonGroup != nil {
				for _, btn := range b.buttonGroup {
					if btn != b && btn.buttonType == ButtonTypeToggle {
						btn.pressed = false
					} else {
						b.pressed = true
					}
				}
			} else {
				b.pressed = !b.pressed
			}
		}
	} else if event == EventActionRelease {
		if b.buttonType == ButtonTypeNormal {
			b.pressed = false
		}
	}
	return true
}
