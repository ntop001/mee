package mee

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func HexToBig(hex string) *big.Int {
	bi := new(big.Int)
	bi.SetString(strings.TrimPrefix(hex, "0x"), 16)
	return bi
}

func BigToHex(bi *big.Int) string {
	return fmt.Sprintf("%064x", bi)
}

func HexToInt64(hex string) int64 {
	i, _ := strconv.ParseInt(strings.TrimPrefix(hex, "0x"), 16, 64)
	return i
}

func Int64ToHex(i64 int64) string {
	return fmt.Sprintf("%064x", i64)
}

func HexToBool(hex string) bool {
	return HexToInt64(hex) > 0
}

func BoolToHex(b bool) string {
	if b {
		return Int64ToHex(1)
	} else {
		return Int64ToHex(0)
	}
}

func HexToStr(hexStr string) string {
	str, _ := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	return string(str)
}

func StrToHex(str string) string {
	_hex := hex.EncodeToString([]byte(str))
	if sz := len(_hex); sz%64 != 0 {
		_hex += strings.Repeat("0", 64-sz%64)
	}
	return _hex
}

func HexToAddress(hex string) string {
	hex = strings.TrimPrefix(hex, "0x")
	if len(hex) > 40 {
		return fmt.Sprintf("0x%s", hex[len(hex)-40:])
	} else {
		return fmt.Sprintf("0x%040s", hex)
	}
}

func AddressToHex(a string) string {
	return fmt.Sprintf("%064s", strings.TrimPrefix(a, "0x"))
}

func HexToBytes(hexStr string) []byte {
	bytes, _ := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	return bytes
}

func BytesToHex(bytes []byte) string {
	_hex := hex.EncodeToString(bytes)
	if sz := len(_hex); sz%64 != 0 {
		_hex += strings.Repeat("0", 64-sz%64)
	}
	return _hex
}