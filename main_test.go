package xuid_test

import (
	"testing"

	"github.com/unhanded/xuid"
)

func TestLite(t *testing.T) {
	id := xuid.NewLite()
	if id == nil {
		t.Fatal("Expected id, got nil")
	}
	t.Logf("%s", id.String())
	t.Logf("%x", id.Bytes())
	t.Logf("%d bytes", len(id.Bytes()))
}

func TestPrefixed(t *testing.T) {
	id := xuid.NewPrefixed([]byte{'a', 'b', 'c'})
	if id == nil {
		t.Fatal("Expected id, got nil")
	}
	t.Logf("%s", id.String())
	t.Logf("%x", id.Bytes())
	t.Logf("%d bytes", len(id.Bytes()))
}
