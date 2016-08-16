package util

import "github.com/satori/go.uuid"

// NewUUID creates a new UUID for usage with rpc.
func NewUUID() string {
	return string(uuid.NewV4().String()[0:16])
}
