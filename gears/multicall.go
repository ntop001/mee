package gears

import (
	"fmt"
	"github.com/ntop001/mee"
	"math/big"
)

// Gear: MultiCall - https://www.multicall3.com/deployments
// ETH Address: 0xcA11bde05977b3631167028862bE2a173976CA11

type Call struct {
	Target string
	Data string
}

type MultiCall struct {
	*mee.Web3Client
	target string
}

func NewMultiCall(targetAddress string, web3 *mee.Web3Client) *MultiCall {
	return &MultiCall{ Web3Client: web3, target: targetAddress }
}

func NewMultiCallUrl(targetAddress, rpcUrl string) *MultiCall  {
	return &MultiCall{ Web3Client: mee.NewWeb3Client(rpcUrl), target: targetAddress}
}

// Aggregate
// function aggregate(Call[] calls) (uint256 blockNumber, bytes[] returnData)
func (mc *MultiCall) Aggregate(calls []Call, args ...interface{}) (*big.Int, []string, error) {
	inputs := make([]interface{}, 0)
	for _, call := range calls {
		inputs = append(inputs, []interface{} { call.Target, mee.HexToBytes(call.Data) })
	}
	callData := mee.TmplEncode([]interface{} { inputs }, "([](address,bytes))")
	result, err := mc.Call(mc.target, fmt.Sprintf("0x252dba42%s", callData), args...)
	if err != nil {
		return nil, nil, err
	}
	results := mee.TmplDecode(result, "(uint256,[]bytes)")
	returnData := make([]string, len(calls))
	for i, v := range results[1].([]interface{}) {
		returnData[i] = mee.BytesToHex(v.([]byte))
	}
	return results[0].(*big.Int), returnData, nil
}

// TryAggregate
// function tryAggregate(bool requireSuccess, Call[] calls) (Result[] returnData)
func (mc *MultiCall) TryAggregate(requireSuccess bool, calls []Call, args ...interface{}) ([]string, error) {
	inputs := make([]interface{}, 0)
	for _, call := range calls {
		inputs = append(inputs, []interface{} { call.Target, mee.HexToBytes(call.Data) })
	}
	callData := mee.TmplEncode([]interface{}{ requireSuccess, inputs }, "(bool,[](address,bytes))")
	result, err := mc.Call(mc.target, fmt.Sprintf("0xbce38bd7%s", callData), args...)
	if err != nil {
		return nil, err
	}
	results := mee.TmplDecode(result, "([](bool,bytes))")
	returnData := make([]string, len(calls))
	for i, v := range results[0].([]interface{}) {
		returnData[i] = mee.BytesToHex(v.([]interface{})[1].([]byte))
	}
	return returnData, nil
}






