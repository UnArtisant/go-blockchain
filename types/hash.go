package types

import (
	"crypto/rand"
)

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		println("HashFromBytes: invalid length")
		panic("HashFromBytes: invalid length")
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomSizeBytes(size int) []byte {
	token := make([]byte, size)

	_, err := rand.Read(token)

	if err != nil {
		return nil
	}

	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomSizeBytes(32))
}
