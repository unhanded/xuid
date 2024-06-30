package xuid

type ID interface {
	String() string
	Bytes() []byte
}
