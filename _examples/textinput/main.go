package main

import (
	"git.sr.ht/~rockorager/vaxis"
	"git.sr.ht/~rockorager/vaxis/widgets/textinput"
)

func main() {
	vx, err := vaxis.New(vaxis.Options{})
	if err != nil {
		panic(err)
	}
	defer vx.Close()
	ti := &textinput.Model{
		Prompt: vaxis.Text{
			Content:    "> ",
			Foreground: vaxis.IndexColor(7),
		},
	}
	for ev := range vx.Events() {
		switch ev := ev.(type) {
		case vaxis.Key:
			switch ev.String() {
			case "Ctrl+c":
				return
			}
		}
		vx.Window().Clear()
		ti.Update(ev)
		ti.Draw(vx.Window())
		vx.Render()
	}
}
