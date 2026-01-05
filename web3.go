package mee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"
)

type Web3Client struct {
	url string
}

func NewWeb3Client(rpcEndpoint string) *Web3Client {
	return &Web3Client{ url: rpcEndpoint }
}

func (web3 *Web3Client) EstimateGas(tx *Tx, args ...interface{}) (int64, error) {
	params := []interface{} {
		ToData(tx),
	}
	if len(args) > 0 {
		params = append(params, args...)
	}
	data, err := web3.RpcCall("eth_estimateGas", params)
	if err != nil {
		return 0, err
	}
	var resultStr string
	if err = json.Unmarshal(data, &resultStr); err != nil {
		return 0, err
	}
	return HexToInt64(resultStr), nil
}

func (web3 *Web3Client) Call(target string, callData string, args ...interface{}) (string, error) {
	params := []interface{} {
		&struct { To string `json:"to"`; Data string `json:"data"`} { target, callData },
	}
	if len(args) > 0 {
		params = append(params, args...)
	}
	data, err := web3.RpcCall("eth_call", params)
	if err != nil {
		return "", err
	}
	var resultStr string
	if err = json.Unmarshal(data, &resultStr); err != nil {
		return "", err
	}
	return resultStr, nil
}

func (web3 *Web3Client) GetTxByHash(hash string) (*TxData, error) {
	params := []interface{} {
		hash,
	}
	data, err := web3.RpcCall("eth_getTransactionByHash", params)
	if err != nil {
		return nil, err
	}
	txData := new(TxData)
	if err = json.Unmarshal(data, txData); err != nil {
		return nil, err
	}
	return txData, nil
}

func (web3 *Web3Client) GetTxReceipt(hash string) (*Receipt, error) {
	params := []interface{} {
		hash,
	}
	data, err := web3.RpcCall("eth_getTransactionReceipt", params)
	if err != nil {
		return nil, err
	}
	receipt := new(Receipt)
	if err = json.Unmarshal(data, receipt); err != nil {
		return nil, err
	}
	return receipt, nil
}

func (web3 *Web3Client) GetBlockNumber() (int64, error) {
	data, err := web3.RpcCall("eth_blockNumber", []interface{}{})
	if err != nil {
		return 0, err
	}
	var numberStr string
	if err = json.Unmarshal(data, &numberStr); err != nil {
		return 0, err
	}
	return HexToInt64(numberStr), nil
}

func (web3 *Web3Client) GetBalance(address string, args ...interface{}) (*big.Int, error) {
	params := []interface{} {
		address,
	}
	if len(args) > 0 {
		params = append(params, args...)
	}
	data, err := web3.RpcCall("eth_getBalance", params)
	if err != nil {
		return nil, err
	}
	var resultStr string
	if err = json.Unmarshal(data, &resultStr); err != nil {
		return nil, err
	}
	return HexToBig(resultStr), nil
}

func (web3 *Web3Client) GetLogs(filter *Filter) ([]*Log, error) {
	rawData, err := web3.RpcCall("eth_getLogs", []interface{}{ filter })
	if err != nil {
		return nil, err
	}
	logs := make([]*Log, 0)
	err = json.Unmarshal(rawData, &logs)
	return logs, err
}

func (web3 *Web3Client) RpcCall(method string, params []interface{}) ([]byte, error) {
	body, err := json.Marshal(&rpcBody{
		JsonRpc: "2.0",
		Id: 1,
		Method: method,
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Post(web3.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &rpcResult{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	if result.Error.Code != 0 {
		return nil, fmt.Errorf("code: %d, msg: %s", result.Error.Code, result.Error.Message)
	}
	return result.Result, err
}

var httpClient = &http.Client{
	Timeout: time.Second * 15,
}

type rpcBody struct {
	JsonRpc string `json:"jsonrpc"`
	Id int `json:"id"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
}

type rpcResult struct {
	JsonRpc string `json:"jsonrpc"`
	Id int `json:"id"`
	Result json.RawMessage `json:"result"`
	Error struct{
		Code int `json:"code"`
		Message string `json:"message"`
		Data string `json:"data"`
	} `json:"error"`
}