package cmd

import (
	"errors"
	"github.com/huanghongxun/agenda/service"
	"github.com/spf13/cobra"
)

var (
	loginCommand = cobra.Command{
		Use:   "login",
		Short: "Log in",
		Long:  "Log in Agenda System with your username and password of a registered account",
	}

	loginUsernameP = loginCommand.Flags().StringP("username", "u", "", "username of a registered account")
	loginPasswordP = loginCommand.Flags().StringP("password", "p", "", "password of a registered account")

	logoutCommand = cobra.Command{
		Use:   "logout",
		Short: "Log out",
		Long:  "Log out current user",
	}

	registerCommand = cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		Long:  "Register a new user with username, password, email and phone number",
	}

	registerUsernameP = registerCommand.Flags().StringP("username", "u", "", "username of the user")
	registerPasswordP = registerCommand.Flags().StringP("password", "p", "", "password of the user")
	registerEmailP    = registerCommand.Flags().StringP("email", "e", "", "email of the user")
	registerPhoneP    = registerCommand.Flags().StringP("phone", "t", "", "phone number of the user")

	listUsersCommand = cobra.Command{
		Use:   "list-users",
		Short: "List all registered users",
		Long:  "List all registered users in Agenda System, logged in required",
	}

	deleteUserCommand = cobra.Command{
		Use:   "delete-user",
		Short: "Delete current user",
		Long:  "Delete current user",
	}
)

func login(cmd *cobra.Command, args []string) error {
	if *loginUsernameP == "" {
		return errors.New("-u required")
	}

	if *loginPasswordP == "" {
		return errors.New("-p required")
	}

	return service.Login(*loginUsernameP, *loginPasswordP)
}

func logout(cmd *cobra.Command, args []string) error {
	return service.Logout()
}

func register(cmd *cobra.Command, args []string) error {
	if *registerUsernameP == "" {
		return errors.New("-u required")
	}

	if *registerPasswordP == "" {
		return errors.New("-p required")
	}

	if *registerEmailP == "" {
		return errors.New("-e required")
	}

	if *registerPhoneP == "" {
		return errors.New("-t required")
	}

	return service.Register(*registerUsernameP, *registerPasswordP, *registerEmailP, *registerPhoneP)
}

func listUsers(cmd *cobra.Command, args []string) error {
	return service.ListUsers()
}

func deleteUser(cmd *cobra.Command, args []string) error {
	return service.DeleteCurrentUser()
}

func init() {
	loginCommand.RunE = login
	logoutCommand.RunE = logout
	registerCommand.RunE = register
	listUsersCommand.RunE = listUsers
	deleteUserCommand.RunE = deleteUser
	rootCmd.AddCommand(&loginCommand, &logoutCommand, &registerCommand, &listUsersCommand, &deleteUserCommand)
}
