# MEE

A lightweight web3 client for self-use😄.<br>
Mee means "Make ethereum easy", get rid of the huge Geth library.

Encoder:
```go
    data := []interface{}{
        big.NewInt(1), 
        int64(2),
        true,
        "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
        "sample",
    }
    hex := mee.TmplEncode(data, "(int256,int64,bool,address,bytes32)")
    fmt.Println("get hex:", hex)
```

Decoder:
```go
    data := "0x0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000464617665000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003"
    tmpl := `(bytes,bool,[]uint256)`

    // decode, dave, true, [1,2,3]
    results := mee.TmplDecode(data, tmpl)
    fmt.Println("get result:", results)
```

## Web3Client

A simple web3 api implementation (only methods I frequently use)

```go
    web3 := mee.NewWeb3Client(os.Getenv("WEB3RPC_URL"))
	
	//run
	balance, err := web3.GetBalance("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get balance:", balance)
```

## Gears

Useful contracts wrappers. Lots to do, I'll add more as I need it.

```go
    w3 := gears.NewMultiCallUrl("0xcA11bde05977b3631167028862bE2a173976CA11", os.Getenv("WEB3RPC_URL"))
	
    //run, multiCall
	calls := []gears.Call {
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
```