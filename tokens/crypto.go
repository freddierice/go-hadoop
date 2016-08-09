package tokens

import "crypto/hmac"

// FillTokenPassword takes a TokenProto and a secret key, and fills in the
// password field of the TokenProto.
func Sign(token Token, key []byte) {
	mac := hmac.New(sha128.New, key)
	mac.Write(token.Bytes())
	token.SetPassword(mac.Sum(nil))
}
