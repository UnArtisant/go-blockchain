package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyPAirVerifySuccess(t *testing.T) {
	priKey := GeneratePrivateKey()
	pubKey := priKey.PublicKey()

	msg := []byte("hello world")
	sig, err := priKey.Sign(msg)

	assert.Nil(t, err)

	b := sig.Verify(pubKey, msg)
	assert.True(t, b)
}

func TestKeyPairVerifyFailMsg(t *testing.T) {
	priKey := GeneratePrivateKey()
	pubKey := priKey.PublicKey()

	msg := []byte("hello world")
	sig, err := priKey.Sign(msg)

	msgChanged := []byte("hello")

	assert.Nil(t, err)

	b := sig.Verify(pubKey, msgChanged)
	assert.False(t, b)
}

func TestKeyPairVerifyFailPubkey(t *testing.T) {
	priKey := GeneratePrivateKey()

	sPriKey := GeneratePrivateKey()
	sPubKey := sPriKey.PublicKey()

	msg := []byte("hello world")
	sig, err := priKey.Sign(msg)

	msgChanged := []byte("hello")

	assert.Nil(t, err)

	b := sig.Verify(sPubKey, msgChanged)
	assert.False(t, b)
}
