package main

import (
	"os"
	"os/exec"

	"git.sr.ht/~rockorager/vaxis"
	"git.sr.ht/~rockorager/vaxis/widgets/border"
	"git.sr.ht/~rockorager/vaxis/widgets/term"
)

type vt struct {
	term *term.Model
}

func newTerm() *vt {
	vt := &vt{}
	vt.term = term.New(vt)
	vt.term.Logger = log
	vt.term.Start(exec.Command(os.Getenv("SHELL")))
	return vt
}

func (vt *vt) Update(msg vaxis.Msg) {
	vt.term.Update(msg)
}

func (vt *vt) Draw(win vaxis.Window) {
	vt.term.Draw(border.All(win, 0, 0))
}
