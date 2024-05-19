/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// importWalletCmd represents the importWallet command
var importWalletCmd = &cobra.Command{
	Use:   "importWallet",
	Short: "Imports an existing wallet.",
	Long: `Imports an existing wallet from a given private key.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Importing wallet...")
		wallet,_ := ImportOldWallet(rpc.DevnetRPCEndpoint)
		fmt.Println("Public Key : " + wallet.account.PublicKey.ToBase58())
		ctx := context.Background()
		balance,_ := GetBalance(ctx)
		fmt.Printf("Wallet Balance : %.9f SOL\n", float64(balance)/1e9)
	},
}

func init() {
	rootCmd.AddCommand(importWalletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
