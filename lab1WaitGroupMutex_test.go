package main

import (
	"testing"
)

func TestXVal(t *testing.T){
	var expected = 0
	if expected != x{
		t.Error("x is not 0")
	}
}

func TestIncrement(t *testing.T){
	var expected = 5
	increment()
	if x != expected {
		t.Errorf("increment fail")
	}
}

func TestMainfunction(t *testing.T){
	var expected = 500
	main()
	if x != expected{
		t.Error("Result false")
	}
}
