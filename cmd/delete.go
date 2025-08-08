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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a password entry",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(vault.DefaultVaultFile); err != nil {
			fmt.Println("Create vault before deleting entries with: pass-vault init")
			return
		}

		if len(args) < 1 {
			fmt.Println("Must enter entry name: pass-vault delete \"entry-name-here\"")
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

		_, exists := v.Entries[entryName]
		if !exists {
			fmt.Printf("Entry doesn't exist: %s\n", entryName)
			return
		} else {
			fmt.Printf("Are you sure you want to delete \"%s\"? (Type 'DELETE' to confirm): ", entryName)

			var confirmation string
			fmt.Scanln(&confirmation)

			if confirmation != "DELETE" {
				fmt.Println("Deletion cancelled")
				return
			}
			delete(v.Entries, entryName)

			err = v.Save(string(masterPassword), vault.DefaultVaultFile)
			if err != nil {
				fmt.Printf("Couldn't save vault: %s\n", err)
				return
			}

			fmt.Println("Entry deleted successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
