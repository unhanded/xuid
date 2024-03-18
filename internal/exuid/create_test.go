package exuid_test

import (
	"testing"

	"github.com/unhanded/exuid/internal/exuid"
)

func TestNewRandomExUID(t *testing.T) {
	out, err := exuid.NewRandomExUID([]byte("WO"))
	if err != nil {
		t.Fatalf("Error creating ExUID: %s", err)
	}
	if len(out) != 20 {
		t.Errorf("Expected length of 20, got %d", len(out))
	}
	t.Logf("FULL HEX: %x", out)
	t.Logf("STRING: %s", out.String())
}

func TestNewRandomExUIDWithOversizedPrefix(t *testing.T) {
	_, err := exuid.NewRandomExUID([]byte("TOOLONG"))
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
