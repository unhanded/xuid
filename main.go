package xuid

import (
	"github.com/unhanded/xuid/internal/lite"
	"github.com/unhanded/xuid/internal/prefixed"
	"github.com/unhanded/xuid/internal/xuid"
)

// Creates a new lite id or dies (a.k.a. panics) trying.
func NewLite() xuid.ID {
	return lite.New()
}

// Creates a new prefixed id. Three bytes are used as a prefix.
func NewPrefixed(pfx []byte) xuid.ID {
	id, err := prefixed.New(pfx)
	if err != nil {
		panic(err)
	}
	return id
}
