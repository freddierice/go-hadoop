package dn

import "testing"

func TestClient(t *testing.T) {
	cl, err := Dial("yes-4.gce.cloudera.com:20002")
	if err != nil {
		t.Fail()
	}

	_ = cl
}
