package ethereum

import (
	"os"
	"os/signal"
	"syscall"

	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	chainsCore "github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Service struct {
	listener   *listener
	connection *connection.Connection
	logger     log15.Logger
	sysErr     chan error
	stop       chan<- int
}

const AliceKey = "alice"
const BobKey = "bob"
const CharlieKey = "charlie"
const DaveKey = "dave"
const EveKey = "eve"

var Keys = []string{AliceKey, BobKey, CharlieKey, DaveKey, EveKey}

func NewService(chainsCfg *chainsCore.ChainConfig, logger log15.Logger) (*Service, error) {
	cfg, err := parseChainConfig(chainsCfg)
	if err != nil {
		return nil, err
	}

	kp, err := makeKey(chainsCfg.KeystorePath)
	sk := crypto.FromECDSA(kp.PrivateKey())
	logger.Info("keypairs:", "public", kp.PublicKey(), "private", hexutil.Encode(sk))
	connection := connection.NewConnection(cfg.endpoint, cfg.http, kp, logger, cfg.gasLimit, cfg.maxGasPrice, cfg.gasMultiplier, cfg.egsApiKey, cfg.egsSpeed)
	err = connection.Connect()
	if err != nil {
		return nil, err
	}

	sysErr := make(chan error)
	stop := make(chan int)
	listener := NewListener(connection, cfg, logger, stop, sysErr, nil)
	ricRegistry, err := NewRICRegistry(cfg.ricRegistryContract, connection.Client())
	listener.setContracts(ricRegistry)

	return &Service{
		connection: connection,
		listener:   listener,
		logger:     logger,
		sysErr:     sysErr,
		stop:       stop,
	}, nil
}

func (s *Service) Start() error {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)

	err := s.listener.start()
	if err != nil {
		return err
	}

	// Block here and wait for a signal
	select {
	case err := <-s.sysErr:
		s.logger.Error("FATAL ERROR. Shutting down.", "err", err)
	case <-sigc:
		s.logger.Warn("Interrupt received, shutting down now.")
	}

	close(s.stop)
	if s.connection != nil {
		s.connection.Close()
	}
	return nil
}
