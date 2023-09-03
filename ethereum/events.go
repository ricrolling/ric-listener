// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"os"
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

var demo = `
	{
    "SystemConfigProxy": "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
    "L1ERC721Bridge": "0xc6e7DF5E7b4f2A278906862b61205850344D4e7d",
    "L1CrossDomainMessengerProxy": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "OptimismMintableERC20Factory": "0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE",
    "L2OutputOracleProxy": "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
    "L1CrossDomainMessenger": "0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82",
    "ProxyAdmin": "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
    "OptimismPortalProxy": "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
    "L2OutputOracle": "0x0B306BF915C4d645ff596e518fAf3F9669b97016",
    "SystemConfig": "0x68B1D87F95878fE05B998F19b66F4baba5De1aed",
    "L1ERC721BridgeProxy": "0x610178dA211FEF7D417bC0e6FeD39F05609AD788",
    "DisputeGameFactory": "0x59b670e9fA9D0A427751Af201D676719a970857b",
    "AddressManager": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    "L1StandardBridge": "0x3Aa5ebB10DC797CAC828524e59A333d0A371443c",
    "L1StandardBridgeProxy": "0x0165878A594ca255338adfa4d48449f69242Eb8F",
    "OptimismMintableERC20FactoryProxy": "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318",
    "OptimismPortal": "0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0",
    "DisputeGameFactoryProxy": "0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e"
}`

var lineaToml = `
l1_chain_id = 59140
l2_chain_id = %d
l1_rpc = "https://linea-goerli.infura.io/v3/0ff5be6b14cc428f840236cacc7bec71"

deployer_key = "0x5a55e41dd61c93435e6594c19ff31f3af72f466bc36bc02decaa30538e847e2c"
batcher_account = "0x7730Edfb83212BABe9396064d765a3d5afEc671a"
batcher_key = "5a55e41dd61c93435e6594c19ff31f3af72f466bc36bc02decaa30538e847e2c"
proposer_account = "0x7730Edfb83212BABe9396064d765a3d5afEc671a"
proposer_key = "5a55e41dd61c93435e6594c19ff31f3af72f466bc36bc02decaa30538e847e2c"
admin_account = "0x7730Edfb83212BABe9396064d765a3d5afEc671a"
admin_key = "5a55e41dd61c93435e6594c19ff31f3af72f466bc36bc02decaa30538e847e2c"
p2p_sequencer_account = "0x7730Edfb83212BABe9396064d765a3d5afEc671a"
p2p_sequencer_key = "5a55e41dd61c93435e6594c19ff31f3af72f466bc36bc02decaa30538e847e2c"`

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (l *listener) writeConfigToml(path string, chainId *big.Int) error {
	l.log.Info("writing config toml", "path", path)
	if err := os.WriteFile(path, []byte(fmt.Sprintf(lineaToml, chainId)), 0777); err != nil {
		return err
	}
	return nil
}

func (l *listener) handleRicRegistryRollUpQueued(chainId *big.Int) error {
	l.log.Info("Handle queued rollup request.", "chainId", chainId)
	interrupt := make(chan int)

	workdir, err := os.MkdirTemp("/tmp", "ric")
	if err != nil {
		l.log.Error("error creating temp dir", "err", err)
		return err
	}
	addressesFile := filepath.Join(workdir, "addresses.json")

	go func() {
		defer close(interrupt)
		err = os.WriteFile(addressesFile, []byte(demo), 0644)
		if err != nil {
			l.log.Error("error writing addresses file", "err", err)
			return
		}
		l.log.Info("Launching rollup", "chainId", chainId)
		time.Sleep(60 * time.Second)
		// shellCmd := fmt.Sprintf("cd /Users/jinsuk/Code/simple-op-stack-rollup/ && python roll.py --name=%s --preset=production --config=./%s l2 &", name, name+".toml")
		// cmd := exec.Command("/bin/zsh", "-c", shellCmd)
		// l.log.Info("Starting roll up service", "cmd", cmd.String())
		// cmd.Stderr = nil
		// output, err := cmd.CombinedOutput()
		// l.log.Info("cmd run output:", "output", string(output))
		// if err != nil {
		// 	l.log.Error("error running rollup script", "err", err)
		// 	return
		// }
		l.log.Info("finished running rollup script")
	}()

	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-l.stop:
				return
			case <-interrupt:
				// temporary for demo.
				l.log.Warn("rollup service exited")
				return
			case <-timer.C:
				l.log.Info("querying", "path", addressesFile)
				if _, err := os.Stat(addressesFile); err == nil {
					l.log.Info("addresses file found, reading...")
					data, err := os.ReadFile(addressesFile)
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

					l.log.Info("payload built", "payload", common.Bytes2Hex(payload))
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
