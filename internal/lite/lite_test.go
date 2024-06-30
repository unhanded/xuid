package lite_test

import (
	"io"
	"testing"
	"time"

	"github.com/unhanded/xuid/internal/lite"
)

func TestLuidString(t *testing.T) {
	id := lite.New()

	t.Log(id.String())
}

func TestliteBytes(t *testing.T) {
	id := lite.New()

	b := id.Bytes()
	if len(b) != 8 {
		t.Errorf("Expected 8 bytes, got %d", len(b))
	}
	t.Logf("%x", b)
}

func TestliteMega(t *testing.T) {
	startTime := time.Now().UnixNano()

	for range 1000000 {
		lite := lite.New()
		b := lite.Bytes()
		io.Discard.Write(b)
	}

	finishTime := time.Now().UnixNano()
	took := finishTime - startTime
	t.Logf("Average %d nanoseconds per generated LUID", took/1000000)
}
