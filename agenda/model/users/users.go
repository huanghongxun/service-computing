package users

import (
	"errors"
	"fmt"
	"github.com/huanghongxun/agenda/model"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

var userStorage = model.Storage{Path: "users.json"}
var users []User

func init() {
	err := userStorage.Load(&users)
	if os.IsNotExist(err) {
		users = []User{}
	} else if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "Unable to load user data from users.json: %s", err.Error()); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func Save() error {
	return userStorage.Save(users)
}

func Add(user *User) (ok bool) {
	_, exists := FindByUsername(user.Username)
	if exists {
		return false
	}
	users = append(users, *user)
	return true
}

func DeleteByUsername(username string) error {
	for i, user := range users {
		if user.Username == username {
			// Remove ith item of users
			users[len(users)-1], users[i] = users[i], users[len(users)-1]
			users = users[:len(users)-1]
			return nil
		}
	}
	return errors.New("User " + username + " not found")
}

func FindByUsername(username string) (user *User, exists bool) {
	for _, user := range users {
		if user.Username == username {
			return &user, true
		}
	}
	return nil, false
}

func ListUsers() []User {
	return users
}
