package main

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/font"
	"github.com/aarzilli/nucular/style"
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"github.com/ryosms/pixela-desktop/internal/login"
	_ "github.com/ryosms/pixela-desktop/statik"
	"io/ioutil"
)

//go:generate statik -f -include=*.ttf -src=fonts
func main() {
	const scale float64 = 1.5
	wnd := nucular.NewMasterWindow(0, "Pixela Desktop", login.UpdateView)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, scale))
	ft := makeFont(scale)
	if ft != nil {
		wnd.Style().Font = *ft
	}
	wnd.Main()
}

func makeFont(scale float64) *font.Face {
	const fontPath string = "/ipaexg00401/ipaexg.ttf"
	const fontSize int = 14

	statikFs, err := fs.New()
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return nil
	}
	f, err := statikFs.Open(fontPath)
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return nil
	}
	defer func() { _ = f.Close() }()

	fontBin, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return nil
	}
	fontFace, err := font.NewFace(fontBin, int(float64(fontSize)*scale))
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return nil
	}
	return &fontFace
}
