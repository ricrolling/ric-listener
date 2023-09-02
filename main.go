package main

import (
	chainsCore "github.com/ChainSafe/chainbridge-utils/core"
	log "github.com/ChainSafe/log15"
	ricEth "github.com/ricrolling/ric-listener/ethereum"
)

func parseConfig(cfgPath string) *chainsCore.ChainConfig {
	cfg := &chainsCore.ChainConfig{}
	return cfg
}

func main() {
	cfg := parseConfig("config.json")
	logger := log.Root().New("chain", cfg.Name)

	s, err := ricEth.NewService(cfg, logger)

	if err != nil {
		logger.Error("failed to create service", "err", err)
		return
	}
	s.Start()
}
