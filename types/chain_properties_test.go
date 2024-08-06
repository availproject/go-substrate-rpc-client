// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

import (
	"testing"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/test_utils"
)

var testChainProperties1 = ChainProperties{}

var testChainProperties2 = ChainProperties{IsEthereum: false, SS58Format: 42,
	TokenDecimals: 18, TokenSymbol: "AVAIL"}

func TestChainProperties_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testChainProperties1)
	AssertRoundtrip(t, testChainProperties2)
}

func TestChainProperties_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{Input: testChainProperties1, Expected: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{Input: testChainProperties2, Expected: []byte{0x00, 0x2a, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x14, 0x41, 0x56, 0x41, 0x49, 0x4c}},
	})
}

func TestChainProperties_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{Input: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Expected: testChainProperties1},
		{Input: []byte{0x00, 0x2a, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x14, 0x41, 0x56, 0x41, 0x49, 0x4c}, Expected: testChainProperties2},
	})
}
