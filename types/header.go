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

package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

type Tuple struct {
	Start  U32
	Offset U32
}
type AppId UCompact
type Value UCompact

type DataLookupIndexItem struct {
	AppId UCompact `json:"appId"`
	Start UCompact `json:"start"`
}
type DataLookup struct {
	Size  UCompact              `json:"size"`
	Index []DataLookupIndexItem `json:"index"`
}

type KateCommitment struct {
	Rows       UCompact `json:"rows"`
	Cols       UCompact `json:"cols"`
	DataRoot   Hash     `json:"dataRoot"`
	Commitment []U8     `json:"commitment"`
}
type KateCommitmentV2 struct {
	Rows       Value `json:"rows"`
	Cols       Value `json:"cols"`
	DataRoot   Hash  `json:"dataRoot"`
	Commitment []U8  `json:"commitment"`
}

type V1HeaderExtension struct {
	Commitment KateCommitment `json:"commitment"`
	AppLookup  DataLookup     `json:"appLookup"`
}
type V2HeaderExtension struct {
	Commitment KateCommitment `json:"commitment"`
	AppLookup  DataLookup     `json:"app_lookup"`
}
type VTHeaderExtension struct {
	NewField   []U8           `json:"newField"`
	Commitment KateCommitment `json:"commitment"`
	AppLookup  DataLookup     `json:"app_lookup"`
}

type HeaderExtensionEnum struct {
	V1 V1HeaderExtension `json:"V1"`
	// V2    V2HeaderExtension `json:"V2"`
	// VTest VTHeaderExtension `json:"VTest"`
}

type HeaderExtension struct {
	V1 V1HeaderExtension `json:"HeaderExtension"`
}

type Header struct {
	ParentHash     Hash                `json:"parentHash"`
	Number         BlockNumber         `json:"number"`
	StateRoot      Hash                `json:"stateRoot"`
	ExtrinsicsRoot Hash                `json:"extrinsicsRoot"`
	Digest         Digest              `json:"digest"`
	Extension      HeaderExtensionEnum `json:"extension"`
}

type BlockNumber U32

// UnmarshalJSON fills BlockNumber with the JSON encoded byte array given by bz
func (b *BlockNumber) UnmarshalJSON(bz []byte) error {
	fmt.Println("umnarshal block number")
	var tmp string
	if err := json.Unmarshal(bz, &tmp); err != nil {
		return err
	}

	s := strings.TrimPrefix(tmp, "0x")

	p, err := strconv.ParseUint(s, 16, 32)
	*b = BlockNumber(p)
	return err
}

// MarshalJSON returns a JSON encoded byte array of BlockNumber
func (b BlockNumber) MarshalJSON() ([]byte, error) {
	fmt.Println("marshal block number")
	s := strconv.FormatUint(uint64(b), 16)
	return json.Marshal(s)
}

// Encode implements encoding for BlockNumber, which just unwraps the bytes of BlockNumber
func (b BlockNumber) Encode(encoder scale.Encoder) error {
	fmt.Println("encode block number")
	return encoder.EncodeUintCompact(*big.NewInt(0).SetUint64(uint64(b)))
}

// Decode implements decoding for BlockNumber, which just wraps the bytes in BlockNumber
func (b *BlockNumber) Decode(decoder scale.Decoder) error {
	fmt.Println("decode block number")
	u, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}
	*b = BlockNumber(u.Uint64())
	return err
}

func (a AppId) Decode(decoder scale.Decoder) error {
	fmt.Println("decode app id")
	u := UCompact(a)
	return u.Decode(decoder)
}

func (a AppId) Encode(encoder scale.Encoder) error {
	fmt.Println("encode app id")
	u := UCompact(a)
	return u.Encode(encoder)
}

func (m HeaderExtension) Encode(encoder scale.Encoder) error {
	var err, err1 error

	err = encoder.PushByte(0)

	if err != nil {
		return err
	}
	err1 = encoder.Encode(m.V1)
	if err1 != nil {
		return err1
	}
	return nil
}

// func (a *Value) Decode(decoder scale.Decoder) error {
// 	fmt.Println("decode value")
// 	u := UCompact(*a)
// 	return u.Decode(decoder)
// 	// ui, err := decoder.DecodeUintCompact()
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// *a = Value(*ui)
// 	// return nil
// }

// func (a Value) Encode(encoder scale.Encoder) error {
// 	fmt.Println("encode value")
// 	u := UCompact(a)
// 	return u.Encode(encoder)
// 	// err := encoder.EncodeUintCompact(big.Int(a))
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// return nil
// }

// func (v *Value) UnmarshalJSON(data []byte) error {
// 	fmt.Println("unmarshal value")
// 	var tmp string
// 	if err := json.Unmarshal(data, &tmp); err != nil {
// 		return err
// 	}
// 	s := strings.TrimPrefix(tmp, "0x")

// 	value := new(big.Int)
// 	if _, ok := value.SetString(s, 16); !ok {
// 		return fmt.Errorf("failed to parse UCompact value")
// 	}
// 	fmt.Println("value", value)
// 	*v = Value(NewUCompact(value))
// 	return nil
// }

//	func (v Value) MarshalJSON() ([]byte, error) {
//		u := UCompact(v)
//		// i := u.Int64()
//		// s := strconv.FormatUint(uint64(i), 16)
//		return u.MarshalJSON()
//		// i := big.Int(v)
//		// return json.Marshal(i.Int64())
//	}
//
// MarshalJSON implements the json.Marshaler interface for HeaderExtension.
