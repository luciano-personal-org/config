package config

import (
	"os"
	"testing"
)

func TestConfigImpl_Get(t *testing.T) {
	// Set up environment variable for testing
	const testKey = "TEST_KEY"
	const testValue = "test_value"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey) // Clean up after test

	// Initialize configImpl
	config := &configImpl{}

	// Test Get method
	got := config.Get(testKey)
	if got != testValue {
		t.Errorf("config.Get(%q) = %q, want %q", testKey, got, testValue)
	}
}

func TestNew(t *testing.T) {
	// Create a temporary .env file for testing
	file, err := os.CreateTemp("", "*.env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up after test

	// Write test environment variable to file
	const testKey = "NEW_TEST_KEY"
	const testValue = "new_test_value"
	_, err = file.WriteString(testKey + "=" + testValue + "\n")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	// Load configuration from the temporary .env file
	config := New(file.Name())
	if config == nil {
		t.Fatal("New() returned nil, want non-nil Config")
	}

	// Test if the environment variable is correctly loaded
	got := config.Get(testKey)
	if got != testValue {
		t.Errorf("config.Get(%q) = %q, want %q", testKey, got, testValue)
	}
}
