/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// getBalanceNewCmd represents the getBalanceNew command
var getBalanceNewCmd = &cobra.Command{
	Use:   "getBalanceNew",
	Short: "Fetch the balance of the new wallet.",
	Long: `This command can be used to fetch the balance of the newly created wallet in the 'data' file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching balance of the new wallet...")
		ctx := context.Background()
		balance, _ := getBalanceNew(ctx)
		fmt.Printf("Wallet Balance : %.9f SOL\n", float64(balance)/1e9)
	},
}

func init() {
	rootCmd.AddCommand(getBalanceNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBalanceNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBalanceNewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
