package mee

import (
	"fmt"
	"math/big"
	"testing"
)

func TestHexToBig(t *testing.T) {
	bi := HexToBig("0x00000000000000000000000000000000000058daae9fced885082f685d6a8f41")
	fmt.Println("get bi:", bi)
}

func TestBigToHex(t *testing.T) {
	bi, _ := big.NewInt(0).SetString("1802177828137172433825747910168385", 10)
	hex := BigToHex(bi)
	fmt.Println("get hex:", hex)
}

func TestHexToInt64(t *testing.T) {
	i64 := HexToInt64("0x0000000000000000000000000000000000000000000000000000000000000080")
	fmt.Println("get bi:", i64)
}

func TestInt64ToHex(t *testing.T) {
	hex := Int64ToHex(128)
	fmt.Println("get hex:", hex)
}

func TestHexToStr(t *testing.T) {
	str := HexToStr("0x73616d706c65")
	fmt.Println("get str:", str)
}

func TestStrToHex(t *testing.T) {
	hex := StrToHex("sample")
	fmt.Println("get hex:", hex)
}

func TestHexToAddress(t *testing.T) {
	address := HexToAddress("0x58daae9fced885082f685d6a8f41")
	fmt.Println("get address:", address)
}

func TestAddressToHex(t *testing.T) {
	hex := AddressToHex("0x00000000000058daae9fced885082f685d6a8f41")
	fmt.Println("get hex:", hex)
}

func TestHexToBytes(t *testing.T) {
	bytes := HexToBytes("0x06fdde03")
	fmt.Println("get bytes:", bytes) //[6 253 222 3]
}

func TestBytesToHex(t *testing.T) {
	hex := BytesToHex([]byte{ 6, 253, 222, 3 })
	fmt.Println("get hex:", hex)
}