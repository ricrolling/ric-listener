package ethereum

import (
	"os"
	"os/signal"
	"syscall"

	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	chainsCore "github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/log15"
)

type Service struct {
	listener   *listener
	connection *connection.Connection
	logger     log15.Logger
	sysErr     chan error
	stop       chan<- int
}

func NewService(chainsCfg *chainsCore.ChainConfig, logger log15.Logger) (*Service, error) {
	cfg, err := parseChainConfig(chainsCfg)
	if err != nil {
		return nil, err
	}

	// TODO: configure chain id correctly
	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainsCfg.Insecure)
	kp, _ := kpI.(*secp256k1.Keypair)

	connection := connection.NewConnection(cfg.endpoint, cfg.http, kp, logger, cfg.gasLimit, cfg.maxGasPrice, cfg.gasMultiplier, cfg.egsApiKey, cfg.egsSpeed)
	err = connection.Connect()
	if err != nil {
		return nil, err
	}

	sysErr := make(chan error)
	stop := make(chan int)
	listener := NewListener(connection, cfg, logger, stop, sysErr, nil)

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
