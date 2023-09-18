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

type DataLookup struct {
	Size  U32      `json:"size"`
	Index [][2]U32 `json:"index"`
}

type KateCommitment struct {
	Rows       U16  `json:"rows"`
	Cols       U16  `json:"cols"`
	DataRoot   Hash `json:"dataRoot"`
	Commitment []U8 `json:"commitment"`
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
	ParentHash     Hash     `json:"parentHash"`
	Number         UCompact `json:"number"`
	StateRoot      Hash     `json:"stateRoot"`
	ExtrinsicsRoot Hash     `json:"extrinsicsRoot"`
	Digest         Digest   `json:"digest"`
	//Extension      HeaderExtensionEnum `json:"extension"`
}
