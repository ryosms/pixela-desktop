package config

import (
	"path/filepath"
	"testing"
)

func testdataDir() string {
	return filepath.FromSlash("../../testdata/internal/config")
}

func TestLoadConfig(t *testing.T) {
	dataDir := testdataDir()
	t.Run("LoadConfig must be able to load config file", func(t *testing.T) {
		configPath := filepath.Join(dataDir, "test_config.toml")
		config, err := LoadConfig(configPath)
		if err != nil {
			t.Fatal(err)
		}
		if len(config.Credentials) != 1 {
			t.Errorf("LoadConfig can't load config!: %+v", config.Credentials)
			return
		}
		cred := config.Credentials[0]
		if cred.Username != "test user" {
			t.Errorf("Username: expected is 'test user' but actual is '%s'", cred.Username)
		}
		if cred.Token != "test token" {
			t.Errorf("Token: expected is 'test token' but actual is '%s'", cred.Token)
		}
	})

}
