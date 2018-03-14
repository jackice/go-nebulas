// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package nvm

import (
	"github.com/nebulasio/go-nebulas/core"
	"github.com/nebulasio/go-nebulas/core/state"
	"github.com/nebulasio/go-nebulas/util/byteutils"
)

const (
	// DefaultLimitsOfTotalMemorySize default limits of total memory size
	DefaultLimitsOfTotalMemorySize uint64 = 40 * 1000 * 1000
)

// Block interface breaks cycle import dependency and hides unused services.
type Block interface {
	Coinbase() *core.Address
	Hash() byteutils.Hash
	Height() uint64
	GetTransaction(hash byteutils.Hash) (*core.Transaction, error)
	RecordEvent(txHash byteutils.Hash, topic, data string) error
}

// SerializableAccount serializable account state
type SerializableAccount struct {
	Nonce   uint64 `json:"nonce"`
	Balance string `json:"balance"`
}

// SerializableBlock serializable block
type SerializableBlock struct {
	Coinbase string `json:"coinbase"`
	Hash     string `json:"hash"`
	Height   uint64 `json:"height"`
}

// SerializableTransaction serializable transaction
type SerializableTransaction struct {
	Hash      string `json:"hash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Nonce     uint64 `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	GasPrice  string `json:"gasPrice"`
	GasLimit  string `json:"gasLimit"`
}

// Context nvm engine context
type Context struct {
	block    Block
	tx       *core.Transaction //ToDo add interface limit
	owner    state.Account
	contract state.Account      //ToDo add interface limit
	state    state.AccountState //ToDo add interface limit name->WorldState
}

// NewContext create a engine context
func NewContext(block Block, tx *core.Transaction, owner state.Account, contract state.Account, state state.AccountState) *Context {
	ctx := &Context{
		block:    block,
		tx:       tx,
		owner:    owner,
		contract: contract,
		state:    state,
	}
	return ctx
}

func toSerializableAccount(acc state.Account) *SerializableAccount {
	sAcc := &SerializableAccount{
		Nonce:   acc.Nonce(),
		Balance: acc.Balance().String(),
	}
	return sAcc
}

func toSerializableBlock(block Block) *SerializableBlock {
	sBlock := &SerializableBlock{
		Coinbase: block.Coinbase().String(),
		Hash:     block.Hash().String(),
		Height:   block.Height(),
	}
	return sBlock
}

func toSerializableTransaction(tx *core.Transaction) *SerializableTransaction {
	sTx := &SerializableTransaction{
		From:      tx.From().String(),
		To:        tx.To().String(),
		Value:     tx.Value().String(),
		Timestamp: tx.Timestamp(),
		Nonce:     tx.Nonce(),
		Hash:      tx.Hash().String(),
		GasPrice:  tx.GasPrice().String(),
		GasLimit:  tx.GasLimit().String(),
	}
	return sTx
}
