// Package config /**
package config

import (
	"bytes"
	"encoding/json"
	"os"
)

var Version = "dev-dirty"

type Log struct {
	LogLevel   int    `json:"LogLevel"`
	LogFileDir string `json:"LogFileDir"`
}

type Backend struct {
	Host    string `json:"Host"`
	RpcPort uint16 `json:"RpcPort"`
}

type BscChain struct {
	Rpc string `json:"Rpc"`
}

type CarrierConfig struct {
	Backend  Backend  `json:"Backend"`
	Logger   Log      `json:"Logger"`
	BscChain BscChain `json:"BscChain"`
}

var CarrierConf CarrierConfig

func init() {
	file, err := os.ReadFile(ConfigName)
	if err != nil {
		panic("config  file error:" + err.Error())
	}
	// Remove the UTF-8 Byte Order Mark
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))

	config := CarrierConfig{}
	if err := json.Unmarshal(file, &config); err != nil {
		panic("unmarshal json config err:" + err.Error())
	}

	CarrierConf = config
}
