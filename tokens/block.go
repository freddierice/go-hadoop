package tokens

import (
	"bytes"
	"time"
)

type BlockAccessMode int

const (
	READ BlockAccessMode = iota
	WRITE
	COPY
	REPLACE
)

type BlockToken struct {
	ExpiryDate  int
	KeyId       int
	UserId      string
	BlockPoolId string
	BlockId     int
	Modes       map[BlockAccessMode]bool
	Password    []byte
	Service     string
}

// Kind implements the Token interface
func (bt *BlockToken) Kind() string {
	return "HDFS_BLOCK_TOKEN"
}

// Password implements the Token interface.
func (bt *BlockToken) Password() []byte {
	return bt.Password
}

// SetPassword implements the Token interface.
func (bt *BlockToken) SetPassword(b []byte) {
	bt.Password = b
}

// Bytes implements the Token interface.
func (bt *BlockToken) Bytes() []byte {
	buf := &bytes.Buffer{}

	hadoop.WriteVInt(buf, bt.ExpiryDate)
	hadoop.WriteVInt(buf, bt.KeyId)
	hadoop.WriteString(buf, bt.UserId)
	hadoop.WriteString(buf, bt.BlockPoolId)
	hadoop.WriteVInt(buf, bt.BlockId)
	hadoop.WriteVInt(buf, len(bt.Modes))
	for _, mode := range bt.Modes {
		hadoop.WriteVInt(buf, mode)
	}

	return buf.Bytes()
}

// NewBlockToken creates a new block token that expires after d amount of time,
// then signs it with the signingKey.
func NewBlockToken(d time.Duration, keyId int, userId string, poolId string,
	blockId int, signingKey []byte, modes ...BlockAccessMode) *BlockToken {
	bt := &BlockToken{
		ExpiryDate:  int(time.Now().Add(d).Unix() * 1000),
		KeyId:       keyId,
		UserId:      userId,
		BlockPoolId: poolId,
		BlockId:     blockId,
	}

	for _, m := range modes {
		bt.Modes[m] = true
	}

	Sign(bt, signingKey)

	return bt
}
