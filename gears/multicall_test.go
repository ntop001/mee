package gears

import (
	"fmt"
	"github.com/ntop001/mee"
	"math/big"
	"os"
	"strings"
	"testing"
)

func TestMultiCall_MultiCall(t *testing.T) {
	w3 := NewMultiCallUrl("0xcA11bde05977b3631167028862bE2a173976CA11", os.Getenv("WEB3RPC_URL"))

	//run, multiCall
	calls := []Call {
		// WETH - decimals()
		{ Target: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", Data: "0x313ce567" },
		// USDc - decimals()
		{ Target: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", Data: "0x313ce567" },
	}
	bn, results, err := w3.Aggregate(calls)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get blockNumber:", bn)

	// decode data from multiCall, parse uint8 as int64
	decimals := make([]int64, len(calls))
	for i, hex := range results {
		decimals[i] = mee.HexToInt64(hex)
	}
	fmt.Println("get results:", decimals)
}

func TestMultiCall_MultiCall1(t *testing.T) {
	w3 := NewMultiCallUrl("0xcA11bde05977b3631167028862bE2a173976CA11", os.Getenv("WEB3RPC_URL"))
	address := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	fnBody := mee.TmplEncode([]interface{}{ address }, `(address)`)

	//run, multiCall
	calls := []Call {
		// WETH - balanceOf(address)
		{ Target: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", Data: fmt.Sprintf("0x70a08231%s", fnBody) },
		// MultiCall.getEthBalance(address)
		{ Target: "0xcA11bde05977b3631167028862bE2a173976CA11", Data: fmt.Sprintf("0x4d2301cc%s", fnBody) },
	}
	bn, results, err := w3.Aggregate(calls)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get blockNumber:", bn)

	// decode data from multiCall, parse uint8 as int64
	decimals := make([]*big.Int, len(calls))
	for i, hex := range results {
		decimals[i] = mee.HexToBig(hex)
	}
	fmt.Println("get weth balance:", decimals[0])
	fmt.Println("get eth balance:", decimals[1])
}

func TestMultiCall_TryMultiCall(t *testing.T) {
	w3 := NewMultiCallUrl("0xcA11bde05977b3631167028862bE2a173976CA11", os.Getenv("WEB3RPC_URL"))

	//run, multiCall
	calls := []Call {
		// WETH - decimals()
		{ Target: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", Data: "0x313ce567" },
		// USDc - decimals()
		{ Target: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", Data: "0x313ce567" },
	}
	results, err := w3.TryAggregate(false, calls)
	if err != nil {
		t.Fatal(err)
	}
	// decode data from multiCall, parse uint8 as int64
	decimals := make([]int64, len(calls))
	for i, hex := range results {
		decimals[i] = mee.HexToInt64(hex)
	}
	fmt.Println("get results:", decimals)
}


func init() {
	// Load .env from root if has one
	envs, err := getEnv("../.env")
	if err != nil {
		return
	}
	for k, v := range envs {
		_ = os.Setenv(k, v)
	}
}

func getEnv(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	envs := make(map[string]string)
	for _, v := range lines {
		if kv := strings.Split(v, "="); len(kv) == 2  {
			envs[kv[0]] = strings.Trim(kv[1], "\"")
		}
	}
	return envs, nil
}