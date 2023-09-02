// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (l *listener) handleRicRegistryRollUpRequested(chainId *big.Int, requester common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return l.ricRegistryContract.QueueRollup(l.conn.Opts(), chainId)
}
