package app

import (
	"os"
	"testing"
)

// Unit tests for LoadConfig.
// These tests verify that LoadConfig correctly loads and validates environment variables,
// applies default values, and returns appropriate errors for missing required config.

// TestLoadConfig_ValidConfig ensures LoadConfig returns a valid Config struct when all env vars are set.
func TestLoadConfig_ValidConfig(t *testing.T) {
	os.Setenv("DB_URL", "postgres://user:pass@localhost:5432/db")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("LOG_LEVEL", "debug")
	defer os.Clearenv()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if cfg.DBURL != "postgres://user:pass@localhost:5432/db" {
		t.Errorf("unexpected DBURL: %s", cfg.DBURL)
	}
	if cfg.ServerPort != "8080" {
		t.Errorf("unexpected ServerPort: %s", cfg.ServerPort)
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("unexpected LogLevel: %s", cfg.LogLevel)
	}
}

// TestLoadConfig_MissingDBURL checks that LoadConfig returns an error if DB_URL is missing.
func TestLoadConfig_MissingDBURL(t *testing.T) {
	os.Unsetenv("DB_URL")
	os.Setenv("SERVER_PORT", "8080")
	defer os.Clearenv()

	_, err := LoadConfig()
	if err == nil || err.Error() != "DB_URL is required" {
		t.Errorf("expected DB_URL is required error, got %v", err)
	}
}

// TestLoadConfig_MissingServerPort checks that LoadConfig returns an error if SERVER_PORT is missing.
func TestLoadConfig_MissingServerPort(t *testing.T) {
	os.Setenv("DB_URL", "postgres://user:pass@localhost:5432/db")
	os.Unsetenv("SERVER_PORT")
	defer os.Clearenv()

	_, err := LoadConfig()
	if err == nil || err.Error() != "SERVER_PORT is required" {
		t.Errorf("expected SERVER_PORT is required error, got %v", err)
	}
}

// TestLoadConfig_DefaultLogLevel ensures that LoadConfig sets LogLevel to "info" if LOG_LEVEL is not set.
func TestLoadConfig_DefaultLogLevel(t *testing.T) {
	os.Setenv("DB_URL", "postgres://user:pass@localhost:5432/db")
	os.Setenv("SERVER_PORT", "8080")
	os.Unsetenv("LOG_LEVEL")
	defer os.Clearenv()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("expected default LogLevel 'info', got %s", cfg.LogLevel)
	}
}
