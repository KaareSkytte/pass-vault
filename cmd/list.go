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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all entries in the vault",

	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(vault.DefaultVaultFile); err != nil {
			fmt.Println("Create vault before adding entries with: pass-vault init")
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

		if len(v.Entries) < 1 {
			fmt.Println("No entries yet")
			fmt.Println("Add entries with: pass-vault add \"entry-name-here\"")
			return
		}

		fmt.Printf("Found %d entries:\n", len(v.Entries))
		for _, entry := range v.Entries {
			fmt.Println(entry.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
