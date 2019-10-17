package service

import (
	"errors"
	"fmt"
	"github.com/huanghongxun/agenda/model/session"
	"github.com/huanghongxun/agenda/model/users"
	"github.com/olekukonko/tablewriter"
	"os"
)

func requireLoggedIn() (username string, err error) {
	username, exists := session.GetCurrentUser()
	if !exists {
		return "", errors.New("not logged in")
	}
	return username, nil
}

func Login(username, password string) error {
	user, exists := users.FindByUsername(username)
	if !exists || user.Password != password {
		return errors.New("username or password is corrupt")
	}
	return session.Login(username)
}

func Logout() error {
	if _, err := requireLoggedIn(); err != nil {
		return err
	}
	return session.Logout()
}

func Register(username, password, email, phone string) error {
	if !users.Add(&users.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}) {
		return errors.New("username is conflict")
	} else {
		if err := users.Save(); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(os.Stderr, "Successfully registered"); err != nil {
			return err
		}
		return nil
	}
}

func ListUsers() error {
	_, err := requireLoggedIn()
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User", "Email", "Phone Number"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, user := range users.ListUsers() {
		table.Append([]string{user.Username, user.Email, user.Phone})
	}
	table.Render()
	return nil
}

func DeleteCurrentUser() error {
	username, err := requireLoggedIn()
	if err != nil {
		return err
	}
	err = users.DeleteByUsername(username)
	if err != nil {
		return err
	}
	err = users.Save()
	if err != nil {
		return err
	}
	return session.Logout()
}
