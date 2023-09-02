package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/ChainBridge/chains"
	chainsEth "github.com/ChainSafe/ChainBridge/chains/ethereum"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

var NotEnoughStake = errors.New("RICRegistry: provider does not have enough stake")
var AlreadyClaimed = errors.New("RICRegistry: rollup not in REQUESTED status or timeout not reached")
var ChainIDNotExist = errors.New("RICRegistry: chainID does not exist")

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("listener block polling failed")

type listener struct {
	cfg                 Config
	conn                chainsEth.Connection
	router              chains.Router
	ricRegistryContract *RICRegistry
	log                 log15.Logger
	stop                <-chan int
	sysErr              chan<- error // Reports fatal error to core
	latestBlock         metrics.LatestBlock
	metrics             *metrics.ChainMetrics
	blockConfirmations  *big.Int
}

// NewListener creates and returns a listener
func NewListener(conn chainsEth.Connection, cfg *Config, log log15.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {

	return &listener{
		cfg:                *cfg,
		conn:               conn,
		log:                log,
		stop:               stop,
		sysErr:             sysErr,
		latestBlock:        metrics.LatestBlock{LastUpdated: time.Now()},
		metrics:            m,
		blockConfirmations: cfg.blockConfirmations,
	}
}

// setContracts sets the listener with the appropriate contracts
func (l *listener) setContracts(ric *RICRegistry) {
	l.ricRegistryContract = ric
}

// sets the router
func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start registers all subscriptions provided by the config
func (l *listener) start() error {
	l.log.Debug("Starting listener...")

	go func() {
		err := l.pollBlocks()
		if err != nil {
			l.log.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	var currentBlock = l.cfg.startBlock
	l.log.Info("Polling Blocks...", "block", currentBlock)

	var retry = BlockRetryLimit
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded")
				l.sysErr <- ErrFatalPolling
				return nil
			}

			latestBlock, err := l.conn.LatestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			if l.metrics != nil {
				l.metrics.LatestKnownBlock.Set(float64(latestBlock.Int64()))
			}

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.blockConfirmations) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			// Parse out events
			err = l.getRicRollupRequestedEvent(currentBlock)
			if err != nil {
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				retry--
				continue
			}

			// Write to block store. Not a critical operation, no need to retry
			if l.metrics != nil {
				l.metrics.BlocksProcessed.Inc()
				l.metrics.LatestProcessedBlock.Set(float64(latestBlock.Int64()))
			}

			l.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			l.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

func (l *listener) getRicRollupRequestedEvent(latestBlock *big.Int) error {
	l.log.Debug("Querying for RIC Rollup Requested event", "block", latestBlock)
	query := buildQuery(l.cfg.ricRegistryContract,
		"rollupRequested(uint256,address,uint256)", latestBlock, latestBlock)
	logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	for _, log := range logs {
		chain := log.Topics[1].Big()
		addr := common.HexToAddress(log.Topics[2].Hex())
		ts := log.Topics[3].Big()

		tx, err := l.handleRicRegistryRollUpRequested(chain, addr, ts)

		if err == NotEnoughStake {
			l.log.Error("Not enough stake", "chain", chain, "addr", addr, "ts", ts, "tx", tx)
			return err
		} else if err == AlreadyClaimed {
			l.log.Error("Already claimed", "chain", chain, "addr", addr, "ts", ts, "tx", tx)
			return err
		} else if err == ChainIDNotExist {
			l.log.Error("Chain ID does not exist", "chain", chain, "addr", addr, "ts", ts, "tx", tx)
			return err
		} else {
			return err
		}
	}

	// TODO: spin it up.
	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig utils.EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
