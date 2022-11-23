package core

import (
	"blockchain/types"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader_Encode_DecodeBinary(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: 123456789,
		Height:    100,
		Nonce:     123456789,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buff))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buff))
	assert.Equal(t, h, hDecode)
}
