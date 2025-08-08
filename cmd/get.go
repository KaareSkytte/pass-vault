/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"

	"github.com/kaareskytte/pass-vault/internal/vault"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a password entry from the vault",

	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(vault.DefaultVaultFile); err != nil {
			fmt.Println("Create vault before getting entries with: pass-vault init")
			return
		}

		if len(args) < 1 {
			fmt.Println("Must enter entry name: pass-vault get \"entry-name-here\"")
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

		entry, exists := v.Entries[entryName]
		if !exists {
			fmt.Printf("Entry \"%s\" not found\n", entryName)
			return
		}

		showPassword, _ := cmd.Flags().GetBool("show-password")

		fmt.Printf("Entry: %s\n", entry.Name)
		fmt.Printf("Username: %s\n", entry.Username)

		if showPassword {
			fmt.Printf("Password: %s\n", entry.Password)
		} else {
			fmt.Printf("Password: [hidden - use --show-password to reveal]\n")
		}

		if entry.URL != "" {
			fmt.Printf("URL: %s\n", entry.URL)
		}

		if entry.Notes != "" {
			fmt.Printf("Notes: %s\n", entry.Notes)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("show-password", "p", false, "Show the password in output")
}
