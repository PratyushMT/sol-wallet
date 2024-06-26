/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// requestAirdropCmd represents the requestAirdrop command
var requestAirdropCmd = &cobra.Command{
	Use:   "requestAirdrop",
	Short: "Requests airdrop in Solana",
	Long: `Requests airdrop to your public address.`,
	Run: func(cmd *cobra.Command, args []string) {
		wallet,_ := ImportOldWallet(rpc.DevnetRPCEndpoint)
		fmt.Println("Requesting Airdrop to " + wallet.account.PublicKey.ToBase58())
		amount,_ := strconv.ParseUint(args[0],10,64)
		txhash,_ := RequestAirDrop(amount)
		fmt.Println("Airdrop of " + strconv.Itoa(int(amount)) + " SOL successful!\nTransaction hash : " + txhash)
	},
}

func init() {
	rootCmd.AddCommand(requestAirdropCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// requestAirdropCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// requestAirdropCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
