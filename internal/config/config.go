package config

import (
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

type Credential struct {
	Username string `toml:"username"`
	Token    string `toml:"token"`
}

type FontConfig struct {
	Path string `toml:"path"`
	Size int    `toml:"size"`
}

type Config struct {
	Font        FontConfig   `toml:"font"`
	Credentials []Credential `toml:"users"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(configPath, &config)
	return &config, err
}

func (config *Config) Save(configPath string) error {
	err := createParentDir(configPath)
	if err != nil {
		return err
	}

	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	return toml.NewEncoder(f).Encode(config)
}

func DefaultConfigPath() string {
	path, _ := homedir.Expand(filepath.FromSlash("~/.pixela/config"))
	return path
}

func createParentDir(filePath string) error {
	dir := filepath.Dir(filePath)
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}
	return os.MkdirAll(dir, 0600)
}
