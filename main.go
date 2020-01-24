package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
	"github.com/ryosms/pixela-desktop/internal/login"
)

func main() {
	wnd := nucular.NewMasterWindow(0, "Pixela Desktop", login.UpdateView)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 1.5))
	wnd.Main()
}
