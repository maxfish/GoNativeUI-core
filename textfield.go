package gui

import (
	"fmt"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"regexp"
)

const validatorInt = "^[-]?[0-9]*$"
const validatorUInt = "^[0-9]*$"
const validatorFloat = "^[-]?[0-9]*[.]?[0-9]+$"

type TextField struct {
	Widget
	LabelText
	IFocusable
	focused bool

	formatValidator *regexp.Regexp
	inputIsValid    bool
	//unit            string
	//defaultText     string

	// Editing
	committed    bool
	editingText  string
	editingRunes []rune
	cursorPos    int
	//selectionStart int
	//selectionEnd   int
}

func NewTextField(text string) *TextField {
	i := &TextField{}
	widgetInit(i)
	i.text = text
	i.Measure()
	return i
}

func NewIntField(value int, signed bool) *TextField {
	text := fmt.Sprintf("%d", value)
	i := NewTextField(text)
	if signed {
		i.SetValidationFormat(validatorInt)
	} else {
		i.SetValidationFormat(validatorUInt)
	}
	return i
}

func NewFloatField(value float32) *TextField {
	text := fmt.Sprintf("%.2f", value)
	i := NewTextField(text)
	i.SetValidationFormat(validatorFloat)
	return i
}

func (i *TextField) initStyle() {
	t := CurrentGui().Theme()
	i.style = &WidgetStyle{
		Font:             t.TextFont,
		FontSize:         t.TextFontSize,
		TextColor:        t.TextColor,
		BackgroundColor:  utils.TransparentColor,
		Padding:          t.ButtonPadding,
		ContentAlignment: utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter},
	}
}

func (i *TextField) Measure() {
	i.computeContentSize()
	i.measuredWidth = i.contentWidth + i.style.Padding.Left + i.style.Padding.Right
	i.measuredHeight = i.contentHeight + i.style.Padding.Top + i.style.Padding.Bottom
	i.measuredFlex = i.flex
}

func (i *TextField) computeContentSize() {
	textSize := i.style.Font.TextSize(i.style.FontSize, i.text)
	i.contentWidth = textSize.W()
	i.contentHeight = textSize.H()
}

func (i *TextField) Text() string {
	if i.focused {
		return i.editingText
	} else {
		return i.text
	}
}

func (i *TextField) CursorPos() int           { return i.cursorPos }
func (i *TextField) Valid() bool              { return i.inputIsValid }

func (i *TextField) SetValidationFormat(format string) {
	i.formatValidator, _ = regexp.Compile(format)
}

func (i *TextField) OnMouseCursorMoved(x float32, y float32) bool {
	return false
}

func (i *TextField) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if button == 0 && event == EventActionPress {
		if !i.focused {
			i.Parent().RequestFocusFor(i)
			return true
		}

		relativeX := int(x - float32(i.bounds.X+i.style.Padding.Left))
		index := i.style.Font.IndexFromCoords(i.style.FontSize, i.text, relativeX, 0)
		i.cursorPos = index
		return true
	}
	return false
}

func (i *TextField) OnKeyEvent(key Key, action EventAction, modifierKey ModifierKey) bool {
	if action == EventActionPress {
		switch key {
		case KeyUp:
			i.cursorPos = 0
		case KeyDown:
			i.cursorPos = len(i.editingRunes)
		case KeyLeft:
			if i.cursorPos > 0 {
				i.cursorPos--
			}
		case KeyRight:
			if i.cursorPos < len(i.editingRunes) {
				i.cursorPos++
			}
		case KeyDelete:
			if i.cursorPos < len(i.editingRunes) {
				i.editingRunes = append(i.editingRunes[:i.cursorPos], i.editingRunes[i.cursorPos+1:]...)
				i.editingText = string(i.editingRunes)
				i.validateInput()
			}
		case KeyBackspace:
			if i.cursorPos > 0 {
				i.editingRunes = append(i.editingRunes[:i.cursorPos-1], i.editingRunes[i.cursorPos:]...)
				i.cursorPos--
				i.editingText = string(i.editingRunes)
				i.validateInput()
			}
		default:
			return false
		}
	}
	return true
}

func (i *TextField) OnCharEvent(char rune) bool {
	i.editingRunes = append(i.editingRunes[:i.cursorPos], append([]rune{rune(char)}, i.editingRunes[i.cursorPos:]...)...)
	i.cursorPos++
	i.editingText = string(i.editingRunes)
	i.validateInput()
	return true
}

func (i *TextField) FocusGained() {
	i.focused = true
	i.editingRunes = []rune(i.text)
	i.editingText = i.text
	i.validateInput()
}

func (i *TextField) FocusLost() {
	i.commitChanges()
	i.focused = false
}

func (i *TextField) Focused() bool {
	return i.focused
}

func (i *TextField) validateInput() {
	if i.committed || i.formatValidator == nil {
		i.inputIsValid = true
		return
	}
	i.inputIsValid = i.formatValidator.MatchString(i.editingText)
}

func (i *TextField) commitChanges() {
	i.validateInput()
	if i.inputIsValid {
		i.text = i.editingText
	}
	i.editingText = i.editingText[:]
	i.editingRunes = i.editingRunes[:]
}
