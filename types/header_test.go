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
	"encoding/hex"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"testing"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/test_utils"
)

var exampleHeader Header

func init() {

	parent, _ := NewHashFromHexString("0x62079bb733a35bd2a22b317162aac307120b02625fdf861f6d88fc08be791de1")
	stateRoot, _ := NewHashFromHexString("0xa1640c227c7a47e160b49a66f867e0fe12a26a3143c420d61409290fc831f4bf")
	extrinsicRoot, _ := NewHashFromHexString("0x869ca1df4fb4440a46929c8cb222d27097f849cc0db3d4cf33f4da1002978029")
	dataRoot, _ := NewHashFromHexString("0x0000000000000000000000000000000000000000000000000000000000000000")

	preRuntime, _ := hex.DecodeString("030000000011380d0500000000f006817ae13bca7d6c4aa181dfeba9231c561b1c8c1406437628f36b297d1b2d58103553567a1b32d5a2156821890da374ca899aad7ef6e1e361accc3ef8fb055e601d47bbda615d6f9776bb40ecdf4127f4206af1523d2444fafeeb8aa2b003")
	seal, _ := hex.DecodeString("58dea98afbcffcfdb6f066728224d5e2e43f45ec38a257e0eb09f370d24af17286d13e341f4f78721614e9873a91f59169fccd256642195e6ef65e7088b44e83")

	fmt.Printf("%v\n", preRuntime)

	exampleHeader = Header{
		ParentHash: parent,
		//Number:         UCompact(big.NewInt(16777215)),
		StateRoot:      stateRoot,
		ExtrinsicsRoot: extrinsicRoot,
		Digest: Digest{
			//{IsChangesTrieRoot: true,
			//	AsChangesTrieRoot: stateRoot,
			//},
			//
			{IsPreRuntime: true, AsPreRuntime: PreRuntime{ConsensusEngineID: 1111573061, Bytes: preRuntime}},
			{IsSeal: true, AsSeal: Seal{ConsensusEngineID: 1111573061, Bytes: seal}},
		},

		Extension: HeaderExtensionEnum{
			V1: V1HeaderExtension{
				Commitment: KateCommitment{
					Rows:     U16(1),
					Cols:     U16(1),
					DataRoot: dataRoot,
					Commitment: []U8{
						69,
					},
				},
				AppLookup: DataLookup{
					Size: 1,
				},
			},
		},
	}
}

// parent hash

func TestHeader_Encoded(t *testing.T) {
	//val, err := Encode(exampleHeader)
	//if err != nil {
	//	fmt.Errorf("Error %v\n", err)
	//}

	//e := U8(255)
	val, _ := EncodeToHex(exampleHeader.Digest)
	//val1, _ := Encode(exampleHeader.Digest[1])
	//var bn BlockNumber
	//_ = Decode(val1, &bn)
	//fmt.Printf("%+v\n", bn)
	fmt.Printf("%+v\n", val)

	api, err := gsrpc.NewSubstrateAPI("ws://127.0.0.1:9944")
	if err != nil {
		fmt.Printf("cannot get api:%v", err)
	}

	b1, _ := hex.DecodeString("0eebb9a3bc45691069710d554567b1166c09798790dcd23be45bc5f851e7c563")
	h256 := NewH256(b1)
	header, err := api.RPC.Chain.GetHeader(Hash(h256))
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	fmt.Printf("%+v\n", header)

	s, _ := EncodeToHex(header.Digest[0].AsPreRuntime.Bytes)
	b, _ := EncodeToHex(header.Digest[0].AsPreRuntime.ConsensusEngineID)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", b)

	s2, _ := EncodeToHex(header.Digest[1].AsSeal.ConsensusEngineID)
	b2, _ := EncodeToHex(header.Digest[1].AsSeal.Bytes)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", b2)

	final, _ := EncodeToHex(header)

	fmt.Printf("%+v\n", final)
	fmt.Printf("%+v\n", header.Number)
	//assert.Equal(t, "0x62079bb733a35bd2a22b317162aac307120b02625fdf861f6d88fc08be791de11112a1640c227c7a47e160b49a66f867e0fe12a26a3143c420d61409290fc831f4bf869ca1df4fb4440a46929c8cb222d27097f849cc0db3d4cf33f4da1002978029080642414245b501030000000011380d0500000000f006817ae13bca7d6c4aa181dfeba9231c561b1c8c1406437628f36b297d1b2d58103553567a1b32d5a2156821890da374ca899aad7ef6e1e361accc3ef8fb055e601d47bbda615d6f9776bb40ecdf4127f4206af1523d2444fafeeb8aa2b0030542414245010158dea98afbcffcfdb6f066728224d5e2e43f45ec38a257e0eb09f370d24af17286d13e341f4f78721614e9873a91f59169fccd256642195e6ef65e7088b44e830004100000000000000000000000000000000000000000000000000000000000000000810198d4ae4db42d766695fa9ed3a5415460cd497cc497f96955e75d3e39881faa70468cf3683c9196a7b74878319542ecb098d4ae4db42d766695fa9ed3a5415460cd497cc497f96955e75d3e39881faa70468cf3683c9196a7b74878319542ecb00400",
	//	hex)

}

func TestHeader_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{Input: exampleHeader, Expected: 272}})
}

func TestHeader_Encode(t *testing.T) {
	x, err := EncodeToHex(exampleHeader)
	fmt.Println(x)
	if err != nil {
		panic(err)
	}
	AssertEncode(t, []EncodingAssert{
		{Input: exampleHeader, Expected: MustHexDecodeString(x)}, //nolint:lll
	})
}

func TestHeader_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{Input: exampleHeader, Expected: "0x0102030405000000000000000000000000000000000000000000000000000000a802030405060000000000000000000000000000000000000000000000000000000304050607000000000000000000000000000000000000000000000000000000140008040502060700000000000000000000000000000000000000000000000000000000000004090000000c0a0b0c050b0000000c0c0d0e060d0000000c0e0f10100408090a0b0c00000000000000000000000000000000000000000000000000000010010203040100000004000000000100000014010203040520080d0e0f1011000000000000000000000000000000000000000000000000000000100506070801000000040000000001000000"}, //nolint:lll
	})
}

func TestHeader_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{Input: exampleHeader, Other: exampleHeader, Expected: true},
		{Input: exampleHeader, Other: NewBytes(hash64), Expected: false},
		{Input: exampleHeader, Other: NewBool(false), Expected: false},
	})
}
