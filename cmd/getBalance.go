/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// getBalanceCmd represents the getBalance command
var getBalanceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "Fetch the balance",
	Long: `Fetch the balance in SOL for the imported wallet`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching balance for the imported wallet...")
		ctx := context.Background()
		balance, _ := GetBalance(ctx)
		fmt.Printf("Wallet Balance : %.9f SOL\n", float64(balance)/1e9)
	},
}

func init() {
	rootCmd.AddCommand(getBalanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
