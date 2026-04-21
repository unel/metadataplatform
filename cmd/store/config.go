package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultMaxLineBytes = 1 << 20 // 1 МБ
const defaultMaxConnections = 1024

type SocketConfig struct {
	Path string   `yaml:"path"`
	Ops  []string `yaml:"ops"`
}

type StoreConfig struct {
	DBUrl          string         `yaml:"db_url"`
	MaxLineBytes   int            `yaml:"max_line_bytes"`
	MaxConnections int            `yaml:"max_connections"`
	Sockets        []SocketConfig `yaml:"sockets"`
}

type Config struct {
	Store      StoreConfig `yaml:"store"`
	configPath string
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config %s: %w", path, err)
	}

	applyDefaults(&cfg)
	return cfg, nil
}

func applyDefaults(cfg *Config) {
	if cfg.Store.MaxLineBytes <= 0 {
		cfg.Store.MaxLineBytes = defaultMaxLineBytes
	}
	if cfg.Store.MaxConnections <= 0 {
		cfg.Store.MaxConnections = defaultMaxConnections
	}
}
