package gui

type Screen struct {
	BoxContainer

	popupStack        []IContainer
	focusedDescendant IFocusable
}

func NewScreen(orientation BoxOrientation, width int, height int) *Screen {
	s := &Screen{}
	containerInit(s)
	s.id = "screen"
	s.orientation = orientation
	s.children = make([]IWidget, 0, 16)
	s.SetWidth(width)
	s.SetHeight(height)
	return s
}

func (s *Screen) RemoveFocusFrom(child IWidget) {
	focusable, ok := child.(IFocusable)
	if ok && focusable == s.focusedDescendant {
		s.focusedDescendant.FocusLost()
		s.focusedDescendant = nil
	}
}

func (s *Screen) RemoveFocus() {
	if s.focusedDescendant != nil {
		s.focusedDescendant.FocusLost()
		s.focusedDescendant = nil
	}
}

func (s *Screen) RequestFocusFor(child IWidget) {
	focusable, ok := child.(IFocusable)
	if ok {
		if focusable == s.focusedDescendant {
			return
		}
		if s.focusedDescendant != nil {
			previousFocused, _ := s.focusedDescendant.(IFocusable)
			previousFocused.FocusLost()
		}
		s.focusedDescendant = focusable
		focusable.FocusGained()
	}
}

func (s *Screen) OnKeyEvent(key Key, action EventAction, modifierKey ModifierKey) bool {
	// Sends the key events only to the focusedDescendant
	if s.focusedDescendant == nil {
		return false
	}
	return s.focusedDescendant.OnKeyEvent(key, action, modifierKey)
}

func (s *Screen) OnCharEvent(char rune) bool {
	// Sends the key events only to the focusedDescendant
	if s.focusedDescendant != nil {
		return s.focusedDescendant.OnCharEvent(char)
	}
	return false
}
