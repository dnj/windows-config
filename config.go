package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Network *NetworkConfig `json:"network"`
	Users   []*UserConfig  `json:"users"`
}

type NetworkConfig struct {
	Addresses []*NetworkAddressConfig `json:"addresses"`
}

type NetworkAddressConfig struct {
	Interface string `json:"interface"`
	Dhcp      bool   `json:"dhcp"`
	Address   string `json:"address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
}

type UserConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ParseConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
