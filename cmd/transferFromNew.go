/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// transferFromNewCmd represents the transferFromNew command
var transferFromNewCmd = &cobra.Command{
	Use:   "transferFromNew",
	Short: "Transfer SOL from your new wallet to other wallets",
	Long: `Use this to transfer SOL from your newly created wallet to other wallets.
	P.S : Remember to airdrop the new wallet from https://faucet.solana.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Recepient address : " + args[0])
		fmt.Println("Amount to be sent : " + args[1] + " SOL")
		amount,_ := strconv.ParseUint(args[1],10,64)
		txhash,_ := TransferFromNew(args[0], amount)
		fmt.Println("Transaction successful.\nTransaction Hash : " + txhash)
	},
}

func init() {
	rootCmd.AddCommand(transferFromNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferFromNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferFromNewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
