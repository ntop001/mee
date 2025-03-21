package web3

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	url string
}

func NewWeb3Client(rpcEndpoint string) *Client {
	return &Client{ url: rpcEndpoint }
}

func (web3 *Client) Call(target string, callData string) (string, error) {
	params := []interface{} {
		&struct { To string `json:"to"`; Data string `json:"data"`} { target, callData },
		"latest",
	}
	data, err := web3.rpcCall("eth_call", params)
	if err != nil {
		return "", err
	}
	var resultStr string
	if err = json.Unmarshal(data, &resultStr); err != nil {
		return "", err
	}
	return resultStr, nil
}

func (web3 *Client) GetBlockNumber() (int64, error) {
	data, err := web3.rpcCall("eth_blockNumber", []interface{}{})
	if err != nil {
		return 0, err
	}
	var numberStr string
	if err = json.Unmarshal(data, &numberStr); err != nil {
		return 0, err
	}
	num, err := strconv.ParseInt(numberStr[2:], 16, 64)
	return num, err
}

func (web3 *Client) rpcCall(method string, params []interface{}) ([]byte, error) {
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
}