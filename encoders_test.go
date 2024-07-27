package abi_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omnes-tech/abi"
)

func ExampleEncode() {
	addressParam := common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")
	uint256ArrParam := []any{big.NewInt(100), big.NewInt(352)}
	bytesParam := []byte("arbitrary byte array...")
	tupleArrParam := []any{
		[]any{&addressParam, uint256ArrParam, bytesParam},
		[]any{&addressParam, uint256ArrParam, bytesParam},
	}

	encoded, err := abi.Encode(
		[]string{"address", "uint256[]", "bytes", "(address,uint256[],bytes)[]"},
		&addressParam, uint256ArrParam, bytesParam, tupleArrParam,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(common.Bytes2Hex(encoded))

	// Output: 0000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e0000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001400000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e0000000000000000000000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e000000000000000000
}

func ExampleEncode_second() {
	encoded, err := abi.Encode(
		[]string{"bytes32[]"},
		[][]byte{
			[]byte("blablabla"),
			[]byte("blablabla2"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(common.Bytes2Hex(encoded))

	// Output: 00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002626c61626c61626c610000000000000000000000000000000000000000000000626c61626c61626c613200000000000000000000000000000000000000000000
}

func ExampleEncodePacked() {
	addressParam := common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")
	uint256ArrParam := []any{big.NewInt(100), big.NewInt(352)}
	bytesParam := []byte("arbitrary byte array...")
	tupleArrParam := []any{
		[]any{&addressParam, uint256ArrParam, bytesParam},
		[]any{&addressParam, uint256ArrParam, bytesParam},
	}

	encoded, err := abi.EncodePacked(
		[]string{"address", "uint256[]", "bytes", "(address,uint256[],bytes)[]"},
		&addressParam, uint256ArrParam, bytesParam, tupleArrParam,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(common.Bytes2Hex(encoded))

	// Output: 5ff137d4b0fdcd49dca30c7cf57e578a026d27890000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000016061726269747261727920627974652061727261792e2e2e5ff137d4b0fdcd49dca30c7cf57e578a026d27890000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000016061726269747261727920627974652061727261792e2e2e5ff137d4b0fdcd49dca30c7cf57e578a026d27890000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000016061726269747261727920627974652061727261792e2e2e
}

func ExampleEncodeWithSignature() {
	funcSignature := "functionName(address,uint256[],bytes,(address,uint256[],bytes)[])"

	addressParam := common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")
	uint256ArrParam := []any{big.NewInt(100), big.NewInt(352)}
	bytesParam := []byte("arbitrary byte array...")
	tupleArrParam := []any{
		[]any{&addressParam, uint256ArrParam, bytesParam},
		[]any{&addressParam, uint256ArrParam, bytesParam},
	}

	encoded, err := abi.EncodeWithSignature(
		funcSignature,
		&addressParam, uint256ArrParam, bytesParam, tupleArrParam,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(common.Bytes2Hex(encoded))

	// Output: c6210dba0000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e0000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001400000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e0000000000000000000000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e000000000000000000
}

func ExampleEncodeWithSelector() {
	funcSignature := "functionName(address,uint256[],bytes,(address,uint256[],bytes)[])"
	selector := abi.EncodeSignature(funcSignature)

	addressParam := common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")
	uint256ArrParam := []any{big.NewInt(100), big.NewInt(352)}
	bytesParam := []byte("arbitrary byte array...")
	tupleArrParam := []any{
		[]any{&addressParam, uint256ArrParam, bytesParam},
		[]any{&addressParam, uint256ArrParam, bytesParam},
	}

	encoded, err := abi.EncodeWithSelector(
		selector,
		[]string{"address", "uint256[]", "bytes", "(address,uint256[],bytes)[]"},
		&addressParam, uint256ArrParam, bytesParam, tupleArrParam,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(common.Bytes2Hex(encoded))

	// Output: c6210dba0000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e0000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001400000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e0000000000000000000000000000000000000000005ff137d4b0fdcd49dca30c7cf57e578a026d2789000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000001761726269747261727920627974652061727261792e2e2e000000000000000000
}

func ExampleEncodeSignature() {
	funcSignature := "functionName(address,uint256[],bytes,(address,uint256[],bytes)[])"
	selector := abi.EncodeSignature(funcSignature)

	fmt.Println(common.Bytes2Hex(selector))

	// Output: c6210dba
}
