package tokens

import (
	hpcommon "gopkg.in/freddierice/go-hproto.v1/common"
)

// Token is an interface for a type that implements the following functions.
type Token interface {
	// GetBytes returns the token in byte form (so it can be read by a hadoop
	// service that uses the Writable interface).
	Bytes() []byte

	// Kind returns the token's kind.
	Kind() string

	// Service returns the token's service.
	Service() string

	// Password returns the token's password.
	Password() []byte

	// SetPassword is a function to set the token's password
	SetPassword([]byte)
}

// NewTokenProto converts a Token to its protobuf version.
func NewTokenProto(t Token) *hpcommon.TokenProto {
	kind := t.Kind()
	service := t.Service()
	return &hpcommon.TokenProto{
		Identifier: t.Bytes(),
		Password:   t.Password(),
		Kind:       &kind,
		Service:    &service,
	}
}
