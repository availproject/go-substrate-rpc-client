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
					//Rows:     U16(1),
					//Cols:     U16(1),
					DataRoot: dataRoot,
					Commitment: []U8{
						69,
					},
				},
				AppLookup: DataLookup{
					//Size: NewUCompactFromUInt(1),
				},
			},
		},
	}
}

// parent hash

func TestHeader_Encoded(t *testing.T) {

	api, err := gsrpc.NewSubstrateAPI("wss://kate.avail.tools/ws")
	if err != nil {
		fmt.Printf("cannot get api:%v", err)
	}

	b1, _ := hex.DecodeString("f340ebeb6a13913f58babae262d888132c99dfbd7ccc0113a9dbc27b4e7e9bbc")
	h256 := NewH256(b1)
	header, err := api.RPC.Chain.GetHeader(Hash(h256))
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	fmt.Printf("%v\n", header)

	//rows, _ := EncodeToHex(header.Extension.V1.Commitment.Rows)
	//fmt.Printf("ROWS %+v   %+v\n", rows, header.Extension.V1.Commitment.Rows)
	//
	//cols, _ := EncodeToHex(header.Extension.V1.Commitment.Cols)
	//fmt.Printf("COLS %+v  %+v   \n", cols, header.Extension.V1.Commitment.Cols)
	//
	//dataLookUP, _ := EncodeToHex(header.Extension.V1.Commitment.Cols)
	//fmt.Printf("DataLookUP %+v\n", dataLookUP)
	//
	//blockNumber, _ := EncodeToHex(header.Number)
	//digest, _ := EncodeToHex(header.Digest)
	//stateRoot, _ := EncodeToHex(header.StateRoot)
	//extrinsicsRoot, _ := EncodeToHex(header.ExtrinsicsRoot)
	//parentHash, _ := EncodeToHex(header.ParentHash)
	//
	//extension, _ := EncodeToHex(header.Extension)
	//commitment, _ := EncodeToHex(header.Extension.V1.Commitment.Commitment)
	//
	//fmt.Printf("\n")
	//fmt.Printf("Digest %+v\n", digest)
	//fmt.Printf("State Root %+v\n", stateRoot)
	//fmt.Printf("Extrinsics Root %+v\n", extrinsicsRoot)
	//fmt.Printf("Block Number %+v\n", blockNumber)
	//fmt.Printf("Parent Hash %+v\n", parentHash)
	//fmt.Printf("Extension %+v\n", extension)
	//fmt.Printf("Commitment %+v\n", commitment)
	//
	//rows, _ := EncodeToHex(header.Extension.V1.Commitment.Rows)
	//cols, _ := EncodeToHex(header.Extension.V1.Commitment.Cols)
	//
	//fmt.Printf("rows %+v\n", rows)
	//fmt.Printf("cols %+v\n", cols)
	//size, _ := EncodeToHex(header.Extension.V1.AppLookup.Size)
	//index, _ := EncodeToHex(header.Extension.V1.AppLookup.Index)
	//fmt.Printf("Size %+v\n", size)
	//fmt.Printf("Index %+v\n", index)
	//
	////
	////Index, _ := EncodeToHex(header.Extension.V1.AppLookup.Index)
	////size, _ := EncodeToHex(header.Extension.V1.AppLookup.Size)
	////
	//fmt.Printf("\n")
	//
	//fmt.Printf("dataLookup===\n")
	//
	////dataLookup, _ := EncodeToHex(header.Extension.V1.AppLookup)
	//
	////fmt.Printf("dataLookup %+v \n", dataLookup)
	//
	////fmt.Printf("Index %+v  %+v\n", Index, header.Extension.V1.AppLookup.Index)
	////fmt.Printf("Size %+v  %+v\n", size, header.Extension.V1.AppLookup.Size)
	////
	////Commitment, _ := EncodeToHex(header.Extension.V1.Commitment.Commitment)
	////fmt.Printf("Commitment %+v\n", Commitment)
	////
	final, _ := EncodeToHex(header)
	//
	fmt.Printf("%+v\n", final)
	//fmt.Printf("%+v\n", header.Number)

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

// 6860 0044
// 6860 0044

//0x31516bef9a85638ad44629b191f8c2ddc57f4a303edc6c7d1c18342bb0672f65368f1b00c29f0a1e0f35005cfe165f3b4bb75abb992bdc601f4b277a95cd726b8ff62e402849cd29af8ee2dd78635e14e1db009d1fdb1ee80b10896c3ba4195b1839f53b080642414245b5010300000000f2430d0500000000febc484f0f37c41fbfb6d3d5a5e72512318bd046f17c5df1c84c9dccdb89a46b7e4b2aa5d3dfa18d3a28d9b0a0777d8187621fe07ae741b5491d90e63fb8040e03a4f7f411a83272644e59e0ffdfd5a61f20a036f673b7f2e802534a759d9a0005424142450101b42cbf49bd343d8ee50b9475d59812c1a005c4826f3d8e385635fc979f802308c2c2f31d71545427347f69dd6b9ef5dff5e11dc95c63706ebb7350da997c06860100100000000000000000000000000000000000000000000000000000000000000000008101b47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298fb47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298f0000
//0x31516bef9a85638ad44629b191f8c2ddc57f4a303edc6c7d1c18342bb0672f65368f1b00c29f0a1e0f35005cfe165f3b4bb75abb992bdc601f4b277a95cd726b8ff62e402849cd29af8ee2dd78635e14e1db009d1fdb1ee80b10896c3ba4195b1839f53b080642414245b5010300000000f2430d0500000000febc484f0f37c41fbfb6d3d5a5e72512318bd046f17c5df1c84c9dccdb89a46b7e4b2aa5d3dfa18d3a28d9b0a0777d8187621fe07ae741b5491d90e63fb8040e03a4f7f411a83272644e59e0ffdfd5a61f20a036f673b7f2e802534a759d9a0005424142450101b42cbf49bd343d8ee50b9475d59812c1a005c4826f3d8e385635fc979f802308c2c2f31d71545427347f69dd6b9ef5dff5e11dc95c63706ebb7350da997c068600044000000000000000000000000000000000000000000000000000000000000000008101b47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298fb47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298f3c00

//0440
