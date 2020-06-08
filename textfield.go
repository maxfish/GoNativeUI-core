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
	IFocusable
	focused bool
	text    string

	offset int
	//unit            string
	//defaultText     string

	// Editing
	formatValidator *regexp.Regexp
	inputIsValid    bool
	committed       bool
	editingText     string
	editingRunes    []rune
	cursorPos       int
	//selectionStart int
	//selectionEnd   int
}

func NewTextField(text string) *TextField {
	i := &TextField{}
	widgetInit(i)
	i.style = CurrentGui().Theme().TextField
	i.text = text
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
	i.style = t.TextField
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

func (i *TextField) TextOffset() int { return i.offset }
func (i *TextField) CursorPos() int  { return i.cursorPos }
func (i *TextField) Valid() bool     { return i.inputIsValid }

func (i *TextField) SetValidationFormat(format string) {
	i.formatValidator, _ = regexp.Compile(format)
}

func (i *TextField) OnMouseEvent(event MouseEvent) IWidget {
	if event.Type == MouseEventButton {
		if event.Button != MouseButtonLeft {
			return nil
		}
		if event.Action == EventActionPress {
			relativeX := int(event.X - float32(i.style.Padding.Left-i.offset))
			text := i.text
			if i.focused {
				text = i.editingText
			}
			index := i.style.Font.IndexFromCoords(i.style.FontSize, text, relativeX, 0)
			i.setCursorPos(index)

			if !i.focused {
				CurrentGui().RequestFocusFor(i)
			}
			return i
		}
	}

	return nil
}

func (i *TextField) OnKeyEvent(key Key, action EventAction, modifierKey ModifierKey) bool {
	if action == EventActionPress {
		switch key {
		case KeyUp:
			i.setCursorPos(0)
		case KeyDown:
			i.setCursorPos(len(i.editingRunes))
		case KeyLeft:
			if i.cursorPos > 0 {
				i.setCursorPos(i.cursorPos - 1)
			}
		case KeyRight:
			if i.cursorPos < len(i.editingRunes) {
				i.setCursorPos(i.cursorPos + 1)
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
				i.editingText = string(i.editingRunes)
				i.setCursorPos(i.cursorPos - 1)
				i.validateInput()
			}
		case KeyEnter:
			i.commitChanges()
		default:
			return false
		}
	}
	return true
}

func (i *TextField) OnCharEvent(char rune) bool {
	i.editingRunes = append(i.editingRunes[:i.cursorPos], append([]rune{rune(char)}, i.editingRunes[i.cursorPos:]...)...)
	i.editingText = string(i.editingRunes)
	i.setCursorPos(i.cursorPos + 1)
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
	i.setCursorPos(0)
	i.focused = false
}

func (i *TextField) Focused() bool {
	return i.focused
}

func (i *TextField) setCursorPos(pos int) {
	i.cursorPos = pos
	textSizeAtCursor := i.style.Font.TextSize(i.style.FontSize, i.editingText, i.cursorPos)
	if textSizeAtCursor.W() > i.offset+i.InnerBounds().W {
		i.offset = textSizeAtCursor.W() - i.InnerBounds().W + 4
	} else if textSizeAtCursor.W() < i.offset {
		i.offset = utils.MaxI(textSizeAtCursor.W()-i.InnerBounds().W/2, 0)
	}
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
