package main

import "testing"

func TestSum(t *testing.T) {
	u := &User{Name: "re", Email: "test"}

	if testgoRoutine(u) == u.Email {
		t.Fail()
	}

}
