package shredder

import "testing"

func TestNoFile(t *testing.T) {
	fileName := "nonexistentfile.txt"
	result := shred(fileName)
	expected := -1

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}