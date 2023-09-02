package ethereum

import (
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	chainsCore "github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/log15"
)

type Service struct {
	listener   *listener
	connection *connection.Connection
	sysErr     chan<- error
	stop       chan<- int
}

func NewService(chainsCfg *chainsCore.ChainConfig, logger log15.Logger) (*Service, error) {
	cfg, err := parseChainConfig(chainsCfg)
	if err != nil {
		return nil, err
	}

	connection := connection.NewConnection(
		cfg.endpoint, cfg.http, nil, logger, cfg.gasLimit, cfg.maxGasPrice, cfg.gasMultiplier, cfg.egsApiKey, cfg.egsSpeed)
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
		sysErr:     sysErr,
		stop:       stop,
	}, nil
}

func (s *Service) Start() error {
	return s.listener.start()
}
