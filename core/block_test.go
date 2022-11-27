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

func TestBlock_Encode_DecodeBinary(t *testing.T) {

	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: 123456789,
			Height:    100,
			Nonce:     123456789,
		},
		Transactions: nil,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buff))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buff))
	assert.Equal(t, b, bDecode)
}

func Test_BlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: 123456789,
			Height:    100,
			Nonce:     123456789,
		},
		Transactions: nil,
	}

	h := b.Hash()
	assert.False(t, h.IsZero())

}
