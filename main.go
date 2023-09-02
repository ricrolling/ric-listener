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
	Name     string            `json:"name"`
	Id       uint8             `json:"id"`
	Endpoint string            `json:"endpoint"`
	From     string            `json:"from"`
	Insecure bool              `json:"insecure"`
	Opts     map[string]string `json:"opts"`
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

	return &chainsCore.ChainConfig{
		Name:     cfg.Name,
		Id:       msg.ChainId(cfg.Id),
		Endpoint: cfg.Endpoint,
		From:     cfg.From,
		Insecure: cfg.Insecure,
		Opts:     cfg.Opts,
	}, nil
}

func run(c *cli.Context) error {
	cfg, err := readConfig(c.String("config"))
	if err != nil {
		return err
	}

	logger := log.Root().New("ric-listener", cfg.Name)
	s, err := ricEth.NewService(cfg, logger)
	if err != nil {
		return err
	}
	s.Start()
	return nil
}

func main() {
	app := &cli.App{
		Name:   "ric-listener",
		Usage:  "RIC Listener",
		Flags:  flags,
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Error("Failed to start", "err", err)
	}
}
