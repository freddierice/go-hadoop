package tokens

import (
	"bytes"
	"io"
	"time"

	"gopkg.in/freddierice/go-hadoop.v2/util"
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
	return ""
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

	util.WriteVInt(buf, bt.ExpiryDate)
	util.WriteVInt(buf, bt.KeyId)
	util.WriteString(buf, bt.UserId)
	util.WriteString(buf, bt.BlockPoolId)
	util.WriteVInt(buf, bt.BlockId)
	util.WriteVInt(buf, len(bt.Modes))
	for mode, _ := range bt.Modes {
		util.WriteVInt(buf, int(mode))
	}

	return buf.Bytes()
}

// ReadBlockToken builds a BlockToken from a block token identifier. Note that
// the token is incomplete -- to add a password, use the SetPassword function.
func ReadBlockToken(r io.Reader) (*BlockToken, error) {
	bt := &BlockToken{}
	var err error
	var nmodes int

	if bt.ExpiryDate, err = util.ReadVInt(r); err != nil {
		return nil, err
	}
	if bt.KeyId, err = util.ReadVInt(r); err != nil {
		return nil, err
	}
	if bt.UserId, err = util.ReadString(r); err != nil {
		return nil, err
	}
	if bt.BlockPoolId, err = util.ReadString(r); err != nil {
		return nil, err
	}
	if bt.BlockId, err = util.ReadVInt(r); err != nil {
		return nil, err
	}
	if nmodes, err = util.ReadVInt(r); err != nil {
		return nil, err
	}
	bt.Modes = make(map[BlockAccessMode]bool, 0)
	for i := 0; i < nmodes; i++ {
		m, err := util.ReadVInt(r)
		if err != nil {
			return nil, err
		}
		bt.Modes[BlockAccessMode(m)] = true
	}

	return bt, nil
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
