package interfaces

import "github.com/chrishanli/tui-go"

// StyledBox is a Box with an overriden Draw method.
// Embedding a Widget within another allows overriding of some behaviors.
type StyledBox struct {
    Style string
    *tui.Box
}

// Draw decorates the Draw call to the widget with a style.
func (s *StyledBox) Draw(p *tui.Painter) {
    p.WithStyle(s.Style, func(p *tui.Painter) {
        s.Box.Draw(p)
    })
}
