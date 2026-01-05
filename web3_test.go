package mee

import (
	"fmt"
	"math/big"
	"os"
	"testing"
)

func TestNewWeb3Client(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	target := "0xf4d2888d29d722226fafa5d9b24f9164c092421e"
	callData := "0x06fdde03"

	// call func name()
	data , err := web3.Call(target, callData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get result:", string(data))
}

func TestWeb3Client_EstimateGas(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	tx := &Tx{
		From: "0xfe3b557e8fb62b89f4916b721be55ceb828dbd73",
		To: "0x44aa93095d6749a706051658b970b941c72c1d53",
		Value: big.NewInt(1),
	}

	//run
	gas, err := web3.EstimateGas(tx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get gas:", gas)
}

func TestWeb3Client_Call(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	target := "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d"
	callData := "0xe985e9c50000000000000000000000007e7022f8879d88bcc5d288b229737adb4b1f39cb0000000000000000000000001e0049783f008a0085193e00003d00cd54003c71"

	// call func isApprovedForAll()
	data , err := web3.Call(target, callData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get result:", string(data))
}

func TestWeb3Client_GetBlockByNumber(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	// run
	b, err := web3.GetBlockByNumber("latest", false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get block number:", b.Number)
	fmt.Println("get block timestamp:", b.Timestamp)
}

func TestWeb3Client_GetTxByHash(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	hash := "0x090d7fdb3f8c0440667404c4f210a51803877f4e86ab1f3b2748ce1df43aa6e8"

	//run
	txData, err := web3.GetTxByHash(hash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get tx data:", txData)
}

func TestWeb3Client_GetTxReceipt(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	hash := "0x090d7fdb3f8c0440667404c4f210a51803877f4e86ab1f3b2748ce1df43aa6e8"

	//run
	receipt, err := web3.GetTxReceipt(hash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get receipt:", receipt)
	fmt.Println("logs:", len(receipt.Logs))
	for _, v := range receipt.Logs {
		fmt.Println(v)
	}
}

func TestWeb3Client_GetBlockNumber(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	// get block number
	num, err := web3.GetBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get blockNumber:", num)
}

func TestWeb3Client_GetBalance(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"

	//run
	balance, err := web3.GetBalance(address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get balance:", balance)
}

func TestWeb3Client_GetLogs(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	//run
	filter := &Filter{
		BlockHash: "0x9ce6a8d622b0c4bf40f51ab4df8acd33792251b319a9afcc91edbd522a5bba35",
		Topics: []string{ "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c" },
	}
	logs, err := web3.GetLogs(filter)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("logs size:", len(logs))
	for _, v := range logs {
		fmt.Println(v)
	}
}

func TestWeb3Client_rpcCall(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	// get balance of account
	data, err := web3.RpcCall("eth_getBalance", []interface{} { "0xbd13c1365b9985387e8e9571cbcd319f7b23daed", "latest" })
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get result:", string(data))
}

