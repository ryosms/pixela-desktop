package login

import (
	"fmt"
	"github.com/aarzilli/nucular"
)

type loginData struct {
	username string
	token    string

	usernameEditor nucular.TextEditor
	tokenEditor    nucular.TextEditor
}

var login loginData

func UpdateView(w *nucular.Window) {
	w.Row(30).Dynamic(1)
	w.Label("Login", "LT")

	w.Row(30).Ratio(0.2, 0.8)

	w.Label("username:", "RC")
	login.usernameEditor.Flags = nucular.EditField
	login.usernameEditor.Edit(w)
	login.username = string(login.usernameEditor.Buffer)

	w.Label("token:", "RC")
	login.tokenEditor.Flags = nucular.EditField
	login.tokenEditor.PasswordChar = 1
	login.tokenEditor.Edit(w)
	login.token = string(login.tokenEditor.Buffer)

	w.Row(30).Dynamic(1)
	if w.ButtonText("Login") {
		fmt.Printf("username: %s, token: %s\n", login.username, login.token)
	}
}
