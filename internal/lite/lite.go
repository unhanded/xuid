package lite

import (
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/unhanded/xuid/internal/xuid"
)

func New() xuid.ID {
	id := liteUID{
		a: rand.Uint32(),
		b: rand.Uint32(),
	}
	return &id
}

type LiteUID interface {
	String() string
	Bytes() []byte
}

type liteUID struct {
	a uint32
	b uint32
}

func (l liteUID) String() string {
	return fmt.Sprintf("%x-%x", l.a, l.b)
}

func (l liteUID) Bytes() []byte {
	bn := binary.LittleEndian.AppendUint32([]byte{}, l.a)
	return binary.LittleEndian.AppendUint32(bn, l.b)
}
