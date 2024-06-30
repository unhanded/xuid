package prefixed

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
)

type Extended [18]byte

func (e Extended) String() string {
	return string(e[:3]) + "-" + hex.EncodeToString(e[3:9]) + "-" + hex.EncodeToString(e[9:12]) + "-" + hex.EncodeToString(e[12:])
}

func (e Extended) Bytes() []byte {
	return e[:]
}

func New(prefix []byte) (Extended, error) {
	if len(prefix) > 3 {
		return Extended{}, errors.New("prefix must be 3 ASCII characters or less")
	}

	b := GenIdBytes()

	prefix = padPrefix(prefix)
	ex := append(prefix, b...)
	return Extended(ex), nil
}

func GenIdBytes() []byte {
	_t := time.Now().UnixMilli()
	t := binary.BigEndian.AppendUint64([]byte{}, uint64(_t))
	t = t[2:]

	var id = make([]byte, 15)
	r := binary.LittleEndian.AppendUint64([]byte{}, rand.Uint64())
	r = binary.LittleEndian.AppendUint64(r, rand.Uint64())

	for idx := range 15 {
		if idx < 6 {
			id[idx] = t[idx]
		} else {
			id[idx] = r[idx-6]
		}
	}
	return id
}

func padPrefix(prefix []byte) []byte {
	pad := make([]byte, 3-len(prefix))
	for i := range pad {
		pad[i] = 0x5f
	}
	return append(prefix, pad...)
}
