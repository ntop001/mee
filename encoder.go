package web3

import (
	"fmt"
	"math/big"
	"strings"
)

var (
	dynamicPlaceholder = strings.Repeat("0", 64)
)

func TmplEncode(data []interface{}, tmpl string) string {
	args := ParseTmpl(tmpl)
	return AbiEncode(data, args)
}

func AbiEncode(data []interface{}, abi []*Arg) string {
	hex, _ := abiEncodeTuple(data, &Arg{ Type: "root", Meta: abi })
	return hex
}

func abiEncodeTuple(data []interface{}, a *Arg) (string, bool) {
	block := make([]string, 0)
	dynamics := make(map[int]string)
	for i, t := range a.Meta {
		b, d := abiEncodeType(data[i], t)
		if d {
			dynamics[i] = b
			block = append(block, dynamicPlaceholder)
		} else {
			block = append(block, b)
		}
	}
	// resolve dynamics
	offset := calBlockSize(block)
	for i := range a.Meta {
		if b := dynamics[i]; len(b) > 0 {
			block[i] = Int64ToHex(offset)
			block = append(block, b)
			offset += int64(len(b)/2)
		}
	}
	return strings.Join(block, ""), len(dynamics) > 0
}

func abiEncodeArray(data []interface{}, a *Arg) (string, bool) {
	block := make([]string, 0)
	if len(a.Meta) == 0 {
		panic("abi encode array: no subtype")
	}

	// encode array value
	dynamics := make(map[int]string)
	subtype := a.Meta[0]
	for i, v := range data {
		b, d := abiEncodeType(v, subtype)
		if d {
			dynamics[i] = b
			block = append(block, dynamicPlaceholder)
		} else {
			block = append(block, b)
		}
	}

	// resolve dynamic
	offset := calBlockSize(block)
	for i := range data {
		if b := dynamics[i]; len(b) > 0 {
			block[i] = Int64ToHex(offset)
			block = append(block, b)
			offset += int64(len(b)/2)
		}
	}

	// encode sz
	sz := int64(len(data))
	hex := Int64ToHex(sz) + strings.Join(block, "")
	return hex, true
}

func calBlockSize(block []string) int64 {
	sz := 0
	for _, v := range block {
		sz += len(v)
	}
	return int64(sz/2)
}

func abiEncodeType(data interface{}, a *Arg) (string, bool) {
	block := ""
	dynamic := false
	switch a.Type {
	case "int64", "uint64":
		block = Int64ToHex(data.(int64))
	case "int256", "uint256":
		block = BigToHex(data.(*big.Int))
	case "bool":
		block = BoolToHex(data.(bool))
	case "address":
		block = AddressToHex(data.(string))
	case "bytes32":
		block = StrToHex(data.(string))
	case "string", "bytes":
		block, dynamic = abiEncodeString(data.(string))
	case "tuple":
		block, dynamic = abiEncodeTuple(data.([]interface{}), a)
	case "array":
		block, dynamic = abiEncodeArray(data.([]interface{}), a)
	default:
		panic("abi unknown type:"+a.Type)
	}
	return block, dynamic
}

func abiEncodeString(str string) (string, bool) {
	sz := int64(len(str))
	block := Int64ToHex(sz)+StrToHex(str)
	return block, true
}

func HexPrettyPrint(block string) {
	sz := len(block)/64
	i := 0
	for ; i < sz; i++ {
		fmt.Println(block[i*64:i*64+64])
	}
	if sz%64 > 0 {
		fmt.Println(block[i*64:])
	}
}