package main

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
)

var count int

func main() {
	wnd := nucular.NewMasterWindow(0, "Pixela Desktop", updatefn)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 1.5))
	wnd.Main()
}

func updatefn(w *nucular.Window) {
	w.Row(50).Dynamic(1)
	if w.ButtonText(fmt.Sprintf("increment: %d", count)) {
		count++
	}

}
