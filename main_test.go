package main

import "testing"

func TestUser(t *testing.T) {
	result := firstTestFunc()

	if result != "firstTestFunc" {
		t.Errorf("Test() FAILED. Expected %s, got %s", "firstTestFunc", result)
	} else {
		t.Logf("Test() PASSED. Expected %s, got %s", "firstTestFunc", result)
	}
}
