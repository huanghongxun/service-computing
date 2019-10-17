package session

import (
	"fmt"
	"github.com/huanghongxun/agenda/model"
	"os"
)

type Session struct {
	LoggedIn bool   `json:"logged_in"`
	Username string `json:"username"`
}

var sessionStorage = model.Storage{Path: "session.json"}
var session Session

func init() {
	err := sessionStorage.Load(&session)
	if os.IsNotExist(err) {
		session = Session{
			LoggedIn: false,
			Username: "",
		}
	} else if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "Unable to load session data from session.json"); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func GetCurrentUser() (username string, exists bool) {
	return session.Username, session.LoggedIn
}

func Login(username string) error {
	session.Username = username
	session.LoggedIn = true
	return sessionStorage.Save(session)
}

func Logout() error {
	session.Username = ""
	session.LoggedIn = false
	return sessionStorage.Save(session)
}
