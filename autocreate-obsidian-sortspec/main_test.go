package main

import (
	"os"
	"testing"
)

func TestGetParams(t *testing.T) {
	// Arrange
	// Mock command line
	os.Args = []string{
		"cmd",
		"-dry",
		"-recursive",
		"-content", "test content",
		"-ext", "txt",
		"/",
	}
	testParams := Params{}
	// Act
	testParams.getParams()
	// Assertion
	want := true
	got := IsValid(testParams.Filename)
	assertMessage(t, want, got)

}

func IsValid(fp string) bool {
	// Check if file exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}
	return false
}

func assertMessage(t testing.TB, want, got bool) {
	t.Helper()
	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}
