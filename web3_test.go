package mee

import (
	"fmt"
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

func TestWeb3Client_GetBlockNumber(t *testing.T) {
	web3 := NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	// get block number
	num, err := web3.GetBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get blockNumber:", num)
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

