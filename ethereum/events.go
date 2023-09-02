// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (l *listener) handleRicRegistryRollUpRequested(chainId *big.Int, requester common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return l.ricRegistryContract.QueueRollup(l.conn.Opts(), chainId)
}

type L1Addresses struct {
	SystemConfigProxy                 string `json:"SystemConfigProxy" validate:"required"`
	L1ERC721Bridge                    string `json:"L1ERC721Bridge" validate:"required"`
	L1CrossDomainMessengerProxy       string `json:"L1CrossDomainMessengerProxy" validate:"required"`
	OptimismMintableERC20Factory      string `json:"OptimismMintableERC20Factory" validate:"required"`
	L2OutputOracleProxy               string `json:"L2OutputOracleProxy" validate:"required"`
	L1CrossDomainMessenger            string `json:"L1CrossDomainMessenger" validate:"required"`
	ProxyAdmin                        string `json:"ProxyAdmin" validate:"required"`
	OptimismPortalProxy               string `json:"OptimismPortalProxy" validate:"required"`
	L2OutputOracle                    string `json:"L2OutputOracle" validate:"required"`
	SystemConfig                      string `json:"SystemConfig" validate:"required"`
	L1ERC721BridgeProxy               string `json:"L1ERC721BridgeProxy" validate:"required"`
	DisputeGameFactory                string `json:"DisputeGameFactory" validate:"required"`
	AddressManager                    string `json:"AddressManager" validate:"required"`
	L1StandardBridge                  string `json:"L1StandardBridge" validate:"required"`
	L1StandardBridgeProxy             string `json:"L1StandardBridgeProxy" validate:"required"`
	OptimismMintableERC20FactoryProxy string `json:"OptimismMintableERC20FactoryProxy" validate:"required"`
	OptimismPortal                    string `json:"OptimismPortal" validate:"required"`
	DisputeGameFactoryProxy           string `json:"DisputeGameFactoryProxy" validate:"required"`
}

func (l *listener) handleRicRegistryRollUpQueued(chainId *big.Int) error {
	l.log.Info("Handle queued rollup request.", "chainId", chainId)
	workdir, err := os.MkdirTemp("/tmp", "rollup")

	if err != nil {
		l.log.Error("error creating temp dir", "err", err)
		return err
	}

	interrupt := make(chan int)

	go func() {
		shellCmd := fmt.Sprintf("python rollup.py --workdir %s --chain_id %d 2>&1", workdir, chainId.Uint64())
		cmd := exec.Command("/bin/zsh", "-c", shellCmd)
		l.log.Info("Starting roll up service", "cmd", cmd.String())
		output, err := cmd.Output()
		l.log.Info("cmd run output:", "output", string(output))
		if err != nil {
			l.log.Error("error running rollup script", "err", err)
		}
		l.log.Info("finished running rollup script")
		close(interrupt)
	}()

	go func() {
		file := filepath.Join(workdir, "addresses.json")
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-l.stop:
				return
			case <-interrupt:
				// temporary for demo.
				time.Sleep(10 * time.Second)
				l.newChainReqCh <- &newChainReq{
					ChainId:     chainId,
					L1Addresses: make([]byte, 18*20+32),
				}
				l.log.Warn("rollup service exited")
				return
			case <-timer.C:
				if _, err := os.Stat(file); err == nil {
					l.log.Info("addresses file found, reading...")
					data, err := os.ReadFile(file)
					if err != nil {
						l.log.Error("error reading addresses file", "err", err)
					}

					var addresses L1Addresses
					err = json.Unmarshal(data, &addresses)
					if err != nil {
						l.log.Error("error unmarshalling addresses file", "err", err)
					}

					l.log.Info("start building payload...")
					payload := common.LeftPadBytes(big.NewInt(18*20).Bytes(), 32)
					payload = append(payload, common.FromHex(addresses.SystemConfigProxy)...)
					payload = append(payload, common.FromHex(addresses.L1ERC721Bridge)...)
					payload = append(payload, common.FromHex(addresses.L1CrossDomainMessengerProxy)...)
					payload = append(payload, common.FromHex(addresses.OptimismMintableERC20Factory)...)
					payload = append(payload, common.FromHex(addresses.L2OutputOracleProxy)...)
					payload = append(payload, common.FromHex(addresses.L1CrossDomainMessenger)...)
					payload = append(payload, common.FromHex(addresses.ProxyAdmin)...)
					payload = append(payload, common.FromHex(addresses.OptimismPortalProxy)...)
					payload = append(payload, common.FromHex(addresses.L2OutputOracle)...)
					payload = append(payload, common.FromHex(addresses.SystemConfig)...)
					payload = append(payload, common.FromHex(addresses.L1ERC721BridgeProxy)...)
					payload = append(payload, common.FromHex(addresses.DisputeGameFactory)...)
					payload = append(payload, common.FromHex(addresses.AddressManager)...)
					payload = append(payload, common.FromHex(addresses.L1StandardBridge)...)
					payload = append(payload, common.FromHex(addresses.L1StandardBridgeProxy)...)
					payload = append(payload, common.FromHex(addresses.OptimismMintableERC20FactoryProxy)...)
					payload = append(payload, common.FromHex(addresses.OptimismPortal)...)
					payload = append(payload, common.FromHex(addresses.DisputeGameFactoryProxy)...)
					l.log.Info("payload built", "payload", payload)

					l.newChainReqCh <- &newChainReq{
						ChainId:     chainId,
						L1Addresses: payload,
					}
					return
				}
			}
		}
	}()

	return nil
}
