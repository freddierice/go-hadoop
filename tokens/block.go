package tokens

import (
	"bytes"
	"time"

	"gopkg.in/freddierice/go-hadoop.v2"
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
	password    []byte
	service     string
}

// Kind implements the Token interface
func (bt *BlockToken) Kind() string {
	return "HDFS_BLOCK_TOKEN"
}

// Service implements the Token interface.
func (bt *BlockToken) Service() string {
	// TODO: make sure this is correct.
	return "hdfs"
}

// Password implements the Token interface.
func (bt *BlockToken) Password() []byte {
	return bt.password
}

// SetPassword implements the Token interface.
func (bt *BlockToken) SetPassword(b []byte) {
	bt.password = b
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
	for mode, _ := range bt.Modes {
		hadoop.WriteVInt(buf, int(mode))
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
