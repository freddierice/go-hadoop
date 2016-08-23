package tokens

import (
	"bytes"
	"encoding/base64"
	"log"
	"testing"
)

func TestBlockToken(t *testing.T) {
	tokIdentEnc := "igFWtS8+hYxv5wFxAAAABHVzZXIAAAAmQlAtNjA4ODY1NTQtMTcyLjMxLjguMTEwLTE0NzE2MjY1Njk3NzaMQAAEOgEEUkVBRA=="
	tokIdent, _ := base64.StdEncoding.DecodeString(tokIdentEnc)
	tokIdentReader := bytes.NewBuffer(tokIdent)

	tok, err := ReadBlockToken(tokIdentReader)
	if err != nil {
		t.Fatal(tok)
	}

	tok2, err := ReadBlockToken(bytes.NewBuffer(tok.Bytes()))
	if err != nil {
		t.Fatalf("could not write then read token: %v", err)
	}

	log.Print(tok)
	log.Print(tok2)

}
