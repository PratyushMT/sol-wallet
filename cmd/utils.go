package cmd

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"errors"
	"fmt"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

type Wallet struct{
	account types.Account
	c *client.Client
}

func CreateNewWallet(RPCEndpoint string) Wallet {
	//Create a new wallet using types.NewAccount()
	newAccount := types.NewAccount()
	//Store the private key in a file named data
	data := []byte(newAccount.PrivateKey)

   err := os.WriteFile("data", data, 0644)
	if err != nil {
			log.Fatal(err)
	}

   return Wallet{
			newAccount,
		client.NewClient(RPCEndpoint),
	}
}

func ImportOldWallet(RPCEndpoint string) (Wallet, error){
	contents, err := os.ReadFile("key_data")
    if err != nil {
        return Wallet{}, err
    }

    //Trim the square brackets and split the contents by commas
    trimmedContents := strings.Trim(string(contents), "[]")
    numbers := strings.Split(trimmedContents, ",")
    privateKey := make([]byte, len(numbers))

    //Convert each number from string to byte
    for i, numberStr := range numbers {
        number, err := strconv.Atoi(strings.TrimSpace(numberStr))
        if err != nil {
            return Wallet{}, errors.New("failed to parse private key: " + err.Error())
        }
        privateKey[i] = byte(number)
    }

    //Create an account from the private key bytes
    account, err := types.AccountFromBytes(privateKey)
    if err != nil {
        return Wallet{}, err
    }

    //Create a new client and return the wallet
    return Wallet{
        account,
		client.NewClient(RPCEndpoint),
    }, nil
}

//Get the balance in lamports.
func GetBalance(ctx context.Context) (uint64, error){
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)

    balance, err := wallet.c.GetBalance(
        ctx,	//Get the context
        wallet.account.PublicKey.ToBase58(),	//Wallet to fetch balance for
    )
    if err != nil {
        return 0, fmt.Errorf("failed to get balance: %w", err)
    }

    return balance, nil
}

func RequestAirDrop(amount uint64) (string,error){
	//Request for SOL
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
	amount = amount * 1e9
	txhash, err := wallet.c.RequestAirdrop(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
		amount,
	)
	if err != nil{
		return "",err
	}
	return txhash, nil
}

func Transfer(receiver string, amount uint64) (string, error) { 
    //Fetch the most recent blockhash 
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    response, err := wallet.c.GetRecentBlockhash(context.TODO()) 
    if err != nil { 
        return "", err 
    } 

    //Make a transfer message with the latest block hash 
    message := types.NewMessage( 
        wallet.account.PublicKey, //Public key of the transaction signer
        []types.Instruction{
            sysprog.Transfer( 
                wallet.account.PublicKey, //Public key of the transaction sender 
                common.PublicKeyFromString(receiver), //Wallet address of the transaction receiver 
                amount*1e9,                               //Transaction amount in lamport 
            ), 
        }, 
        response.Blockhash, // recent block hash 
    ) 

    //Create a transaction with the message and TX signer 
    tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account}) 
    if err != nil { 
        return "", err 
    } 

    //Send the transaction to the blockchain 
    txhash, err := wallet.c.SendTransaction2(context.TODO(), tx) 
    if err != nil { 
        return "", err 
    } 
    return txhash, nil 
}

func switchToNew(RPCEndpoint string) (Wallet,error){
	content, _ := os.ReadFile("data")
	privateKey := []byte(string(content))
	wallet, err := types.AccountFromBytes(privateKey)

	if err != nil{
		return Wallet{}, err
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func getBalanceNew(ctx context.Context) (uint64, error){
	wallet, _ := switchToNew(rpc.DevnetRPCEndpoint)

    balance, err := wallet.c.GetBalance(
        ctx,	//Get the context
        wallet.account.PublicKey.ToBase58(),	//Wallet to fetch balance for
    )
    if err != nil {
        return 0, fmt.Errorf("failed to get balance: %w", err)
    }

    return balance, nil
}

func TransferFromNew(receiver string, amount uint64) (string, error) { 
    //Fetch the most recent blockhash 
    wallet, _ := switchToNew(rpc.DevnetRPCEndpoint) 
    response, err := wallet.c.GetRecentBlockhash(context.TODO()) 
    if err != nil { 
        return "", err 
    } 

    //Make a transfer message with the latest block hash 
    message := types.NewMessage( 
        wallet.account.PublicKey, //Public key of the transaction signer
        []types.Instruction{
            sysprog.Transfer( 
                wallet.account.PublicKey, //Public key of the transaction sender 
                common.PublicKeyFromString(receiver), //Wallet address of the transaction receiver 
                amount*1e9,                               //Transaction amount in lamport 
            ), 
        }, 
        response.Blockhash, // recent block hash 
    ) 

    //Create a transaction with the message and TX signer 
    tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account}) 
    if err != nil { 
        return "", err 
    } 

    //Send the transaction to the blockchain 
    txhash, err := wallet.c.SendTransaction2(context.TODO(), tx) 
    if err != nil { 
        return "", err 
    } 
    return txhash, nil 
}