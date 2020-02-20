package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/font"
	"github.com/aarzilli/nucular/style"
	"github.com/ryosms/pixela-desktop/internal/config"
	"github.com/ryosms/pixela-desktop/internal/login"
	"io/ioutil"
)

func main() {
	const scale float64 = 1.5
	wnd := nucular.NewMasterWindow(0, "Pixela Desktop", login.UpdateView)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, scale))

	conf, err := config.LoadConfig(config.DefaultConfigPath())
	fontPath := "ipaexg.ttf"
	fontSize := 14
	if err == nil {
		if len(conf.Font.Path) > 0 {
			fontPath = conf.Font.Path
		}
		if conf.Font.Size > 0 {
			fontSize = conf.Font.Size
		}
	}

	ftb, err := ioutil.ReadFile(fontPath)
	if err == nil {
		ft, err := font.NewFace(ftb, int(float64(fontSize)*scale))
		if err == nil {
			wnd.Style().Font = ft
		}
	}
	wnd.Main()
}
