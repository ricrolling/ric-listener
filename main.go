package main

import (
	"encoding/json"
	"os"

	chainsCore "github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
	ricEth "github.com/ricrolling/ric-listener/ethereum"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "config",
		Usage:    "Path to json config `FILE`",
		Required: true,
	},
}

type Config struct {
	Name               string `json:"name"`
	Id                 uint8  `json:"id"`
	Endpoint           string `json:"endpoint"`
	From               string `json:"from"`
	Insecure           bool   `json:"insecure"`
	KeystorePath       string `json:"keystorePath"`
	RicRegistryAddress string `json:"ricRegistry"`
}

func readConfig(jsonCfg string) (*chainsCore.ChainConfig, error) {
	var cfg *Config
	data, err := os.ReadFile(jsonCfg)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	opts := make(map[string]string)
	opts[ricEth.RicRegistryOpt] = cfg.RicRegistryAddress

	return &chainsCore.ChainConfig{
		Name:         cfg.Name,
		Id:           msg.ChainId(cfg.Id),
		Endpoint:     cfg.Endpoint,
		From:         cfg.From,
		KeystorePath: cfg.KeystorePath,
		Insecure:     cfg.Insecure,
		Opts:         opts,
	}, nil
}

func run(c *cli.Context) error {
	cfg, err := readConfig(c.String("config"))
	if err != nil {
		return err
	}

	logger := log.Root().New("ric-service", cfg.Name)
	s, err := ricEth.NewService(cfg, logger)
	if err != nil {
		return err
	}
	s.Start()
	return nil
}

func main() {
	app := &cli.App{
		Name:   "ric-service",
		Usage:  "RIC service",
		Flags:  flags,
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Error("Failed to start", "err", err)
	}
}
