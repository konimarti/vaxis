package button

import (
	"git.sr.ht/~rockorager/vaxis"
	"git.sr.ht/~rockorager/vaxis/widgets/align"
)

type SizePolicy func(vaxis.Window) vaxis.Window

var FixedWidth func(width int) SizePolicy = func(width int) SizePolicy {
	return func(w vaxis.Window) vaxis.Window {
		w.Width = width
		return w
	}
}

var FlexWidth SizePolicy = func(w vaxis.Window) vaxis.Window { return w }

type Model struct {
	label string

	policy SizePolicy
	style  vaxis.Style

	onPress func() interface{}
}

func New(s string, onPress func() interface{}, sp SizePolicy) *Model {
	return &Model{label: s, policy: sp, onPress: onPress}
}

func (m *Model) SetStyle(style vaxis.Style) *Model {
	m.style = style
	return m
}

func (m *Model) Press() interface{} {
	if m.onPress != nil {
		return m.onPress()
	}
	return nil
}

func (m *Model) Update(msg vaxis.Event) {
}

func (m *Model) Draw(win vaxis.Window) {
	w, h := win.Size()
	if w == 0 || h == 0 {
		return
	}
	win.Clear()

	policyWin := m.policy(win)
	button := align.TopMiddle(win, policyWin.Width, 1)

	button.Fill(vaxis.Cell{
		Character: vaxis.Character{
			Grapheme: " ",
			Width:    1,
		},
		Style: vaxis.Style{Background: m.style.Background},
	})

	align.TopMiddle(button, len(m.label), 1).Println(0, vaxis.Segment{Text: m.label, Style: m.style})
}
