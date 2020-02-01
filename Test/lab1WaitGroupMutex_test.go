package main

import (
	// "sync" // TODO: Use sync with WaitGroup and Mutex!
	"testing"
)

func TestXVal(t *testing.T){
	var expected = 0
	if expected != x{
		t.Error("x is not 0")
	}
}
