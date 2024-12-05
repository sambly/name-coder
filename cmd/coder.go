/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/sambly/name-coder/internal"
	"github.com/spf13/cobra"
)

var coderCmd = &cobra.Command{
	Use:   "coder",
	Short: "name-coder coder --name= --password= ",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")

		out, err := internal.Encrypt(name, password)
		if err != nil {
			fmt.Println("Error coder func ")
			return
		}
		fmt.Printf("Name coder = %s", out)
	},
}

func init() {
	rootCmd.AddCommand(coderCmd)

	coderCmd.Flags().String("name", "", "Имя пользователя")
	coderCmd.Flags().String("password", "", "Пароль для шифрования")

	coderCmd.MarkFlagRequired("name")
	coderCmd.MarkFlagRequired("password")
}
