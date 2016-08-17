package dn

import "testing"

func TestClient(t *testing.T) {
	token := "igFWkt3EDIxJWp4RAAAABGhkZnMAAAAoQlAtMjE0NTMzMzU4My0xNzIuMzEuOS4xMTktMTQ2ODI4NDA4NTA0MIxAANwRAQVXUklURQ=="
	cl, err := Dial(token, "hi", "yes-3.gce.cloudera.com:20002")
	if err != nil {
		t.Fatal(err)
	}

	_ = cl
}
