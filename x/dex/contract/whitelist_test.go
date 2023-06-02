package contract_test

import (
	"encoding/hex"
	"testing"

	"github.com/sei-protocol/sei-chain/x/dex/contract"
	"github.com/stretchr/testify/require"
)

func TestGetWasmPrefixes(t *testing.T) {
	wasmWhitelistedPrefixes := contract.GetWasmWhitelistedPrefixes("sei14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sh9m79m")

	wasmPrefixBytes, _ := hex.DecodeString("03" + "ade4a5f5803a439835c636395a8d648dee57b2fc90d98dc17fa887159b69638b")
	require.Equal(t, []byte(wasmWhitelistedPrefixes[0]), wasmPrefixBytes)
}

