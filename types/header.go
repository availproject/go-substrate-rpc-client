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
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"math/big"
	"strconv"
	"strings"
)

type DataLookupItem struct {
	AppId UCompact
	Start UCompact
}

type DataLookup struct {
	Size  UCompact         `json:"size"`
	Index []DataLookupItem `json:"index"`
}

type KateCommitment struct {
	Rows       U16  `json:"rows"`
	Cols       U16  `json:"cols"`
	DataRoot   Hash `json:"dataRoot"`
	Commitment []U8 `json:"commitment"`
}

func (u U16) Encode(encoder scale.Encoder) error {

	var i = big.NewInt(int64(u))

	err := encoder.EncodeUintCompact(*i)
	if err != nil {
		return err
	}
	return nil
}

type V1HeaderExtension struct {
	Commitment KateCommitment `json:"commitment"`
	AppLookup  DataLookup     `json:"appLookup"`
}
type VTHeaderExtension struct {
	NewField   []U8           `json:"newField"`
	Commitment KateCommitment `json:"commitment"`
	AppLookup  DataLookup     `json:"appLookup"`
}
type HeaderExtensionEnum struct {
	V1 V1HeaderExtension `json:"v1"`
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
	s := strconv.FormatUint(uint64(b), 16)
	return json.Marshal(s)
}

// Encode implements encoding for BlockNumber, which just unwraps the bytes of BlockNumber
func (b BlockNumber) Encode(encoder scale.Encoder) error {
	return encoder.EncodeUintCompact(*big.NewInt(0).SetUint64(uint64(b)))
}

// Decode implements decoding for BlockNumber, which just wraps the bytes in BlockNumber
func (b *BlockNumber) Decode(decoder scale.Decoder) error {
	u, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}
	*b = BlockNumber(u.Uint64())
	return err
}

//0x31516bef9a85638ad44629b191f8c2ddc57f4a303edc6c7d1c18342bb0672f65368f1b00c29f0a1e0f35005cfe165f3b4bb75abb992bdc601f4b277a95cd726b8ff62e402849cd29af8ee2dd78635e14e1db009d1fdb1ee80b10896c3ba4195b1839f53b080642414245b5010300000000f2430d0500000000febc484f0f37c41fbfb6d3d5a5e72512318bd046f17c5df1c84c9dccdb89a46b7e4b2aa5d3dfa18d3a28d9b0a0777d8187621fe07ae741b5491d90e63fb8040e03a4f7f411a83272644e59e0ffdfd5a61f20a036f673b7f2e802534a759d9a0005424142450101b42cbf49bd343d8ee50b9475d59812c1a005c4826f3d8e385635fc979f802308c2c2f31d71545427347f69dd6b9ef5dff5e11dc95c63706ebb7350da997c06860 0044   000000000000000000000000000000000000000000000000000000000000000008101b47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298fb47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298f3c00
//0x31516bef9a85638ad44629b191f8c2ddc57f4a303edc6c7d1c18342bb0672f65368f1b00c29f0a1e0f35005cfe165f3b4bb75abb992bdc601f4b277a95cd726b8ff62e402849cd29af8ee2dd78635e14e1db009d1fdb1ee80b10896c3ba4195b1839f53b080642414245b5010300000000f2430d0500000000febc484f0f37c41fbfb6d3d5a5e72512318bd046f17c5df1c84c9dccdb89a46b7e4b2aa5d3dfa18d3a28d9b0a0777d8187621fe07ae741b5491d90e63fb8040e03a4f7f411a83272644e59e0ffdfd5a61f20a036f673b7f2e802534a759d9a0005424142450101b42cbf49bd343d8ee50b9475d59812c1a005c4826f3d8e385635fc979f802308c2c2f31d71545427347f69dd6b9ef5dff5e11dc95c63706ebb7350da997c06860 100100 000000000000000000000000000000000000000000000000000000000000000008101b47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298fb47c7ec08e8069b95d8fa80947b26c04d7ecc945913685659c22bb9d506c070ed9c103c3970336a622e27924e787298f0000
