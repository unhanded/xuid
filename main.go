package xuid

import "github.com/unhanded/xuid/internal/xuid"

// Creates a new xuid or dies (a.k.a. panics) trying.
func New(pfx []byte) xuid.XUID {
	id, err := xuid.New(pfx)
	if err != nil {
		panic(err)
	}
	return id
}
