package prefixed_test

import (
	"testing"

	"github.com/unhanded/xuid/internal/prefixed"
)

func TestGenIdBytes(t *testing.T) {
	b := prefixed.GenIdBytes()
	if len(b) != 15 {
		t.Errorf("Expected 15 bytes, got %d\n", len(b))
	}
	t.Logf("%x", b)
}

func TestRandomextended(t *testing.T) {
	id, err := prefixed.New([]byte{'t', 's', 't'})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", id.String())
	t.Logf("%x\n", id)
	t.Logf("%d bytes\n", len(id))
}

func TestIncorrect(t *testing.T) {
	_, err := prefixed.New([]byte{'f', 'a', 'i', 'l'})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
