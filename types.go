package mee

import (
	"fmt"
	"math/big"
)

type Tx struct {
	From string `json:"from"`
	To string `json:"to"`
	Data string `json:"data"`
	Value *big.Int `json:"value"`

	// optional
	ChainId int64 `json:"chainId"`
	Nonce *big.Int `json:"nonce"`
	Gas *big.Int `json:"gas"`
	Type int64 `json:"type"`
	// gas type=2, only
	MaxFeePerGas *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
}

func ToData(tx *Tx) *TxData {
	txData := &TxData{
		From: tx.From,
		To: tx.To,
		Input: tx.Data,
	}
	if tx.Value != nil {
		txData.Value = fmt.Sprintf("0x%x", tx.Value)
	}
	if tx.ChainId > 0 {
		txData.ChainId = fmt.Sprintf("0x%x", tx.ChainId)
	}
	if tx.Nonce != nil {
		txData.Nonce = fmt.Sprintf("0x%x", tx.Nonce)
	}
	if tx.Gas != nil {
		txData.Gas = fmt.Sprintf("0x%x", tx.Gas)
	}
	if tx.Type > 0 {
		txData.Type = fmt.Sprintf("0x%x", tx.Type)
	}
	if tx.MaxFeePerGas != nil {
		txData.MaxFeePerGas = fmt.Sprintf("0x%x", tx.MaxFeePerGas)
	}
	if tx.MaxPriorityFeePerGas != nil {
		txData.MaxPriorityFeePerGas = fmt.Sprintf("0x%x", tx.MaxPriorityFeePerGas)
	}
	return txData
}

type TxData struct {
	BlockHash string `json:"blockHash,omitempty"`
	BlockNumber string `json:"blockNumber,omitempty"`
	ChainId string `json:"chainId,omitempty"`
	From string `json:"from,omitempty"`
	To string `json:"to,omitempty"`
	Gas string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Type string `json:"type,omitempty"`
	MaxFeePerGas string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	Hash string `json:"hash,omitempty"`
	Input string `json:"input,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	Value string `json:"value,omitempty"`
	V string `json:"v,omitempty"`
	R string `json:"r,omitempty"`
	S string `json:"s,omitempty"`
}

type Receipt struct {
	BlockHash string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	ContractAddress string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	From string `json:"from"`
	To string `json:"to"`
	GasUsed string `json:"gasUsed"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	Logs []*Log `json:"logs"`
	LogsBloom string `json:"logsBloom"`
	Status string `json:"status"`
	TransactionHash string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
	Type string `json:"type"`
}

type Block struct {
	Number string `json:"number"`
	Hash string `json:"hash"`
	MixHash string `json:"mixHash"`
	ParentHash string `json:"parentHash"`
	Nonce string `json:"nonce"`
	Sha3Uncles string `json:"sha3Uncles"`
	LogsBloom string `json:"logsBloom"`
	TransactionsRoot string `json:"transactionsRoot"`
	StateRoot string `json:"stateRoot"`
	ReceiptsRoot string `json:"receiptsRoot"`
	Miner string `json:"miner"`
	Difficulty string `json:"difficulty"`
	TotalDifficulty string `json:"totalDifficulty"`
	ExtraData string `json:"extraData"`
	Size string `json:"size"`
	GasLimit string `json:"gasLimit"`
	GasUsed string `json:"gasUsed"`
	Timestamp string `json:"timestamp"`
	Uncles []string `json:"uncles"`
	Transactions []string `json:"transactions"`
	BaseFeePerGas string `json:"baseFeePerGas"`
	WithdrawalsRoot string `json:"withdrawalsRoot"`
	BlobGasUsed string `json:"blobGasUsed"`
	ExcessBlobGas string `json:"excessBlobGas"`
	ParentBeaconBlockRoot string `json:"parentBeaconBlockRoot"`
}

type Filter struct {
	Address string `json:"address,omitempty"`
	FromBlock string `json:"fromBlock,omitempty"`
	ToBlock string `json:"toBlock,omitempty"`
	BlockHash string `json:"blockHash,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
}

type Log struct {
	Address string `json:"address"`
	BlockHash string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	Data string `json:"data"`
	LogIndex string `json:"logIndex"`
	Removed bool `json:"removed"`
	Topics []string `json:"topics"`
	TransactionHash string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
}