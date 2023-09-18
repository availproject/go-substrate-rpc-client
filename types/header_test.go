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
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/test_utils"
)

var exampleHeader Header

func init() {

	parent, _ := NewHashFromHexString("0x2390716cf47146b869ea093625bce79d6466622275c6e6c7572d7bdc173db88c")
	stateRoot, _ := NewHashFromHexString("0x98e4ebfb4514857a3fb86c6f4aee55d1640f3005ede69ed186218a9870967d5c")
	extrinsicRoot, _ := NewHashFromHexString("0x19e3798ba6d1b7ec0842d76c8f4c7eb25ce39e4e60a42c0350904aa910f82c90")
	dataRoot, _ := NewHashFromHexString("0x0000000000000000000000000000000000000000000000000000000000000000")

	preRuntime, _ := hex.DecodeString("0300000000b3350d050000000036239c4f99b0106087ee43817ed0bd1ddfaf43f4dff20b58582cb18ac1131f688bb33e1a51f4048f83e6ef39d623bc60c9314495801ac8d769073cbaabb2700618842b66414594c456b88d52fec4ff6e6cad68a2c39b5c78307cbd567ca2cf0f")
	seal, _ := hex.DecodeString("18960a95b37451727f8d756d31f1b58f294adcfa3faaf45adb15d169be666e0fc587f290d910e5118b4470a5bd44500735434d052f965353bdd46d4ce18eef8a")

	exampleHeader = Header{
		ParentHash:     parent,
		Number:         NewU32(550),
		StateRoot:      stateRoot,
		ExtrinsicsRoot: extrinsicRoot,
		Digest: Digest{
			{IsPreRuntime: true, AsPreRuntime: PreRuntime{ConsensusEngineID: 1111573061, Bytes: NewBytes(preRuntime)}},
			{IsSeal: true, AsSeal: Seal{ConsensusEngineID: 1111573061, Bytes: NewBytes(seal)}},
		},

		Extension: HeaderExtensionEnum{
			V1: V1HeaderExtension{
				Commitment: KateCommitment{
					Rows:     U16(1),
					Cols:     U16(4),
					DataRoot: dataRoot,
					Commitment: []U8{
						173, 99, 172, 161, 54, 91, 154, 130, 184, 238, 174, 6,
						185, 233, 233, 199, 17, 95, 183, 53, 22, 43, 157, 129,
						237, 94, 99, 21, 196, 218, 156, 88, 52, 137, 182, 181,
						121, 252, 248, 74, 61, 232, 42, 129, 222, 67, 129, 85,
						173, 99, 172, 161, 54, 91, 154, 130, 184, 238, 174, 6,
						185, 233, 233, 199, 17, 95, 183, 53, 22, 43, 157, 129,
						237, 94, 99, 21, 196, 218, 156, 88, 52, 137, 182, 181,
						121, 252, 248, 74, 61, 232, 42, 129, 222, 67, 129, 85,
					},
				},
				AppLookup: DataLookup{
					Size: 1,
				},
			},
		},
	}
}

func TestHeader_Encoded(t *testing.T) {

	//vc, _ := json.Marshal(exampleHeader)
	//fmt.Printf("%s\n", vc)

	s, _ := json.MarshalIndent(exampleHeader, "", "   ")

	fmt.Printf("%s\n", s)

	v, _ := Encode(exampleHeader)

	expected := Header{}
	Decode(v, &expected)

	fmt.Printf("%v\n", expected)

	//for _, a := range v {
	//	fmt.Printf("%v\n", a)
	//}

	//fmt.Printf("%v\n", v)

}

//0x// var (
// 	headerFuzzOpts = digestItemFuzzOpts
// )

// func TestHeader_EncodeDecode(t *testing.T) {
// 	AssertRoundtrip(t, exampleHeader)
// 	AssertRoundTripFuzz[Header](t, 100, headerFuzzOpts...)
// 	AssertDecodeNilData[Header](t)
// 	AssertEncodeEmptyObj[Header](t, 98)
// }

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
