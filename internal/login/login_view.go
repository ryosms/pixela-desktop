package login

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/ryosms/pixela-desktop/internal/graphs"
	"github.com/ryosms/pixela-desktop/pkg/pixela"
	"strings"
)

type loginData struct {
	username string
	token    string

	usernameEditor nucular.TextEditor
	tokenEditor    nucular.TextEditor

	message string
}

var login loginData

func UpdateView(w *nucular.Window) {
	w.Row(30).Dynamic(1)
	w.Label("Login", "LT")

	w.Row(30).Ratio(0.2, 0.8)

	w.Label("username:", "RC")
	login.usernameEditor.Flags = nucular.EditField | nucular.EditIbeamCursor
	login.usernameEditor.Edit(w)
	login.username = string(login.usernameEditor.Buffer)

	w.Label("token:", "RC")
	login.tokenEditor.Flags = nucular.EditField | nucular.EditIbeamCursor
	login.tokenEditor.PasswordChar = '*'
	login.tokenEditor.Edit(w)
	login.token = string(login.tokenEditor.Buffer)

	w.Row(30).Dynamic(1)
	w.Label(login.message, "LC")

	w.Row(30).Dynamic(1)
	if w.ButtonText("Login") {
		if len(strings.TrimSpace(login.username)) == 0 || len(strings.TrimSpace(login.token)) == 0 {
			login.message = "username and token are required."
			return
		}
		graphList, err := pixela.GetGraphDefinitions(login.username, login.token)
		if err != nil {
			fmt.Printf("%+v\n", err)
			login.message = "login failed."
		} else {
			login.message = ""
			graphs.ShowList(w, login.username, graphList)
		}
	}
}
