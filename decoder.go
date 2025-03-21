package web3

import (
	"fmt"
	"strings"
)

type Arg struct {
	Type string // string, int256, tuple, array
	Meta []*Arg
}

// ABI spec: https://docs.soliditylang.org/en/latest/abi-spec.html
// For composable types:
// Array: &Arg { Type: "array", Meta: []*Arg{ Int256 } }
// Tuple: &Arg { Type: "tuple", Meta: []*Arg{ String, Int256..} }
var (
	String  = &Arg{ Type: "string" }
	Address = &Arg{ Type: "address" }
	Int256  = &Arg{ Type: "int256" }
	Uint256 = &Arg{ Type: "uint256" }
	Bool    = &Arg{ Type: "bool" }
	Bytes32 = &Arg{ Type: "bytes32" }
	Bytes   = &Arg{ Type: "bytes" }
)

// TmplDecode decode with a params template
func TmplDecode(hex string, tmpl string) []interface{} {
	args := ParseTmpl(tmpl)
	return AbiDecode(hex, args)
}

func AbiDecode(hex string, abi []*Arg) []interface{} {
	_hex := strings.TrimPrefix(hex, "0x")
	v, _ := abiDecodeTuple(_hex, _hex, &Arg{ Type: "root", Meta: abi })
	return v
}

func abiDecodeTuple(b, block string, a *Arg) ([]interface{}, int) {
	if a.Type != "root" && isDynamic(a) {
		hex64 := b[:64]
		offset := HexToInt64(hex64)
		b = block[offset*2:]
	}
	values := make([]interface{}, 0)
	pos := 0
	for _, t := range a.Meta {
		val, sz := abiDecodeType(b[pos:], b, t)
		values = append(values, val)
		pos += sz
	}
	if isDynamic(a) {
		return values, 64
	} else {
		return values, pos
	}
}

func abiDecodeArray(b string, block string, a *Arg) ([]interface{}, int) {
	offset := HexToInt64(b[:64])
	valBlock := block[offset*2:]
	sz := int(HexToInt64(valBlock[:64]))
	values := make([]interface{}, sz)
	pos := 64
	if len(a.Meta) == 0 {
		panic("abi decode array: no subtype")
	}
	subType := a.Meta[0]
	for i := 0; i < sz; i++ {
		val, szz := abiDecodeType(valBlock[pos:], valBlock[64:], subType)
		values[i] = val
		pos += szz
	}
	return values, 64
}

func abiDecodeType(b, block string, a *Arg) (interface{}, int) {
	hex64 := b[:64]
	var val interface{}
	sz := 64
	switch a.Type {
	case "int256", "uint256", "bool", "address":
		val = hex64
	case "bytes32":
		val = hex64
	case "bytes", "string":
		val = abiDecodeBytes(hex64, block)
	case "tuple":
		val, sz = abiDecodeTuple(b, block, a)
	case "array":
		val, sz = abiDecodeArray(b, block, a)
	default:
		panic("abi unknown type: "+a.Type)
	}
	return val, sz
}

func abiDecodeBytes(hex64, block string) string {
	offset := HexToInt64(hex64)
	valBlock := block[offset*2:]
	sz := HexToInt64(valBlock[:64])
	return valBlock[64: 64+sz*2]
}

//if tuple has dynamic type, it's a dynamic tuple
//which will be encoded as dynamic.
func isDynamic(a *Arg) bool {
	switch a.Type {
	case "array", "string", "bytes":
		return true
	case "tuple":
		for _, v := range a.Meta {
			if ok := isDynamic(v); ok { return true }
		}
	}
	return false
}

func (a Arg) String() string {
	if a.Type == "tuple" {
		return fmt.Sprintf("tuple(%s)", a.formatMeta())
	}
	if a.Type == "array" && len(a.Meta) > 0 {
		return fmt.Sprintf("[]%s", a.Meta[0])
	}
	return a.Type
}

func (a Arg) formatMeta() string {
	if len(a.Meta) == 0 {
		return ""
	}
	subs := a.Meta[0].String()
	for _, v := range a.Meta[1:] {
		subs += ","+v.String()
	}
	return subs
}