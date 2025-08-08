/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"

	"github.com/kaareskytte/pass-vault/internal/vault"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password entry",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(vault.DefaultVaultFile); err != nil {
			fmt.Println("Create vault before adding entries with: pass-vault init")
			return
		}

		if len(args) < 1 {
			fmt.Println("Must enter entry name: pass-vault add \"entry-name-here\"")
			return
		}

		fmt.Print("Enter master password: ")
		masterPassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Invalid password")
			return
		}
		fmt.Println()

		v, err := vault.LoadVault(string(masterPassword), vault.DefaultVaultFile)
		if err != nil {
			fmt.Printf("Couldn't load vault: %v", err)
			return
		}

		entryName := args[0]

		fmt.Print("Enter username: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		username := strings.TrimSpace(scanner.Text())

		if strings.Contains(username, " ") {
			fmt.Println("Error: Username cannot contain spaces")
			return
		}

		fmt.Print("Enter password: ")
		password, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Invalid password")
			return
		}
		fmt.Println()

		entry := vault.Entry{
			ID:       entryName,
			Name:     entryName,
			Username: username,
			Password: string(password),
			Created:  time.Now(),
			Updated:  time.Now(),
		}

		url, _ := cmd.Flags().GetString("url")
		if url != "" {
			entry.URL = url
		}

		note, _ := cmd.Flags().GetString("note")
		if note != "" {
			entry.Notes = note
		}

		err = v.AddEntry(entryName, entry)
		if err != nil {
			fmt.Printf("Couldn't add entry: %v\n", err)
			return
		}

		err = v.Save(string(masterPassword), vault.DefaultVaultFile)
		if err != nil {
			fmt.Printf("Couldn't save entry: %v\n", err)
			return
		}

		fmt.Println("Entry added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("url", "u", "", "Add url for entry")
	addCmd.Flags().StringP("note", "n", "", "Add note for entry")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
