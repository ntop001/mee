package example

import (
	"fmt"
	"github.com/ntop001/mee"
	"os"
	"strings"
	"testing"
)

// Example: send a simple call
func Test_GetDecimal(t *testing.T)  {
	w3 := mee.NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	//run, erc20 decimals
	result, err := w3.Call("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0x313ce567")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get decimals:", mee.HexToInt64(result))
}

// Example: encode input and decode output
// BAYC: 0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d
// func: isApprovedForAll(address owner, address operator) (bool)
// bytes4: 0xe985e9c5
func Test_isApprovedForAll(t *testing.T) {
	w3 := mee.NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	//run, erc721 fetch approve status
	inputs := []interface{} {
		"0x7e7022f8879d88bcc5d288b229737adb4b1f39cb", //owner
		"0x1e0049783f008a0085193e00003d00cd54003c71", //operator
	}
	data := mee.TmplEncode(inputs, "(address,address)")
	result, err := w3.Call("0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d", fmt.Sprintf("0xe985e9c5%s", data))
	if err != nil {
		t.Fatal(err)
	}
	outputs := mee.TmplDecode(result, "(bool)")
	fmt.Println("get approval status:", outputs[0])
}

// Example: a more complex case, get decimal of wETH, USDc
// MultiCall on ETH: 0xcA11bde05977b3631167028862bE2a173976CA11
// Call{ address, bytes }
// function aggregate(Call[] calls) (uint256 blockNumber, bytes[] returnData) - byte4: 0x252dba42
// ERC20: function decimals() (uint8) - byte4:0x313ce567
func Test_MultiCall(t *testing.T) {
	w3 := mee.NewWeb3Client(os.Getenv("WEB3RPC_URL"))

	//run, multiCall
	calls := []interface{} {
		[]interface{}{ "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", mee.HexToBytes("0x313ce567") },
		[]interface{}{ "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", mee.HexToBytes("0x313ce567") },
	}
	inputs := []interface{}{ calls }
	data := fmt.Sprintf("0x252dba42%s", mee.TmplEncode(inputs, "([](address,bytes))"))
	result, err := w3.Call("0xcA11bde05977b3631167028862bE2a173976CA11", data)
	if err != nil {
		t.Fatal(err)
	}
	outputs := mee.TmplDecode(result, "(uint256,[]bytes)")
	fmt.Println("get blockNumber:", outputs[0])

	// decode data from multiCall, parse uint8 as int64
	// Note: returnData is bytes([]byte internally), convert to hex then mee can parse it.
	decimals := make([]int64, len(calls))
	for i, returnData := range outputs[1].([]interface{}) {
		hex := mee.BytesToHex(returnData.([]byte))
		decimals[i] = mee.HexToInt64(hex)
		// Another way (if returnData is not a single value):
		// v := mee.TmplDecode(hex, "(int64)")
		// decimal[i] = v[0]
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