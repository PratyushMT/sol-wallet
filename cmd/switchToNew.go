/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// switchToNewCmd represents the switchToNew command
var switchToNewCmd = &cobra.Command{
	Use:   "switchToNew",
	Short: "Switch to the new wallet.",
	Long: `Switch to the newly created wallet.
		This command imports the wallet from 'data' file which was created using the cli command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Switching to the new wallet...")
		wallet,_ := switchToNew(rpc.DevnetRPCEndpoint)
		fmt.Println("Public Key : " + wallet.account.PublicKey.ToBase58())
		ctx := context.Background()
		balance,_ := getBalanceNew(ctx)
		fmt.Printf("Wallet Balance : %.9f SOL\n", float64(balance)/1e9)
	},
}

func init() {
	rootCmd.AddCommand(switchToNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchToNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchToNewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
