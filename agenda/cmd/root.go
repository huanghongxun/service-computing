package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "CLI App for user and meeting management",
	Long: `Agenda is an application for user and meeting management.

You can register a new user, login or delete some users`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
