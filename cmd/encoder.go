/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/sambly/name-coder/internal"
	"github.com/spf13/cobra"
)

var encoderCmd = &cobra.Command{
	Use:   "encoder",
	Short: "name-coder encoder --name= --password= ",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")

		out, err := internal.Decrypt(name, password)
		if err != nil {
			fmt.Println("Error encoder func ")
			return
		}
		fmt.Printf("Name encoder = %s", out)
	},
}

func init() {
	rootCmd.AddCommand(encoderCmd)

	encoderCmd.Flags().String("name", "", "Имя пользователя")
	encoderCmd.Flags().String("password", "", "Пароль для шифрования")

	encoderCmd.MarkFlagRequired("name")
	encoderCmd.MarkFlagRequired("password")
}
