package web3

import (
	"fmt"
	"testing"
)

func TestParseTmpl(t *testing.T) {
	tmpl := "(int256,string,[](int256,int16),(uint256,string))"

	//parse
	args := ParseTmpl(tmpl)
	fmt.Println("get args", args)
}

func TestParser_getTuple_0(t *testing.T) {
	tmpl := "(uint256, string),"

	//parse
	arg, sz := getTuple(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}

func TestParser_getTuple_1(t *testing.T) {
	tmpl := "(uint256, []string),"

	//parse
	arg, sz := getTuple(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}

func TestParser_getTuple_2(t *testing.T) {
	tmpl := "(uint256, (string, int256)),"

	//parse
	arg, sz := getTuple(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}

func TestParser_getArray_0(t *testing.T) {
	tmpl := "[]uint256,"

	//parse
	arg, sz := getArray(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}

func TestParser_getArray_1(t *testing.T) {
	tmpl := "[](uint256, string),"

	//parse
	arg, sz := getArray(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}

func TestParser_getType_0(t *testing.T) {
	tmpl := "uint256, string"

	//parse
	arg, sz := getType(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}

func TestParser_getType_1(t *testing.T) {
	tmpl := "uint256), string"

	//parse
	arg, sz := getType(tmpl)
	fmt.Println("get result:", arg, "left:", tmpl[sz:])
}