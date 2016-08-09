package tokens

import "gopkg.in/freddierice/go-hadoop.v1/hproto"

// Token is an interface for a type that implements the following functions.
type Token interface {
	// GetBytes returns the token in byte form (so it can be read by a hadoop
	// service that uses the Writable interface).
	Bytes() []byte

	// Kind returns the token's kind.
	Kind() string

	// Password returns the token's password.
	Password() string

	// SetPassword is a function to set the token's password
	SetPassword(string)
}

// NewTokenProto converts a Token to its protobuf version.
func NewTokenProto(t Token) *hproto.TokenProto {
	return &hproto.TokenProto{
		Identifier: t.Bytes(),
		Password:   t.Password(),
		Kind:       t.Kind(),
		Service:    t.Service(),
	}
}
