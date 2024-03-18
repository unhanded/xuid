package exuid

import (
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
)

type ExUID [20]byte

func (e ExUID) String() string {
	return string(e[:4]) + "-" + hex.EncodeToString(e[5:9]) + "-" + hex.EncodeToString(e[10:12]) + "-" + hex.EncodeToString(e[13:15]) + "-" + hex.EncodeToString(e[16:])
}

type ExUIDData struct {
	Prefix ExUIDAsciiPrefix
	Uuid   uuid.UUID
}

func NewRandomExUID(prefix []byte) (ExUID, error) {
	if len(prefix) > 4 {
		return ExUID{}, errors.New("prefix must be 4 ASCII characters or less")
	}
	baseUuid, err := uuid.NewV7()
	if err != nil {
		return ExUID{}, err
	}
	prefix = padPrefix(prefix)
	ex := append(prefix, baseUuid[:]...)
	return ExUID(ex), nil
}

func padPrefix(prefix []byte) []byte {
	pad := make([]byte, 4-len(prefix))
	for i := range pad {
		pad[i] = 0x2f
	}
	return append(prefix, pad...)
}
