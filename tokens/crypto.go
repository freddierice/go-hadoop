package tokens

import (
	"crypto/hmac"
	"crypto/sha1"
)

// FillTokenPassword takes a TokenProto and a secret key, and fills in the
// password field of the TokenProto.
func Sign(token Token, key []byte) {
	mac := hmac.New(sha1.New, key)
	mac.Write(token.Bytes())
	token.SetPassword(mac.Sum(nil))
}
