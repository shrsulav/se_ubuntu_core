package shredder

import "testing"

// test for shredding a file which does not exist
func Test_1(t *testing.T) {
	fileName := "nonexistentfile.txt"
	result := shred(fileName)
	expected := -1

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}

// test for shredding a file which exists and has good file permissions
func Test_2(t *testing.T) {
	fileName := "../test_files/test_file_1.txt"

	result := shred(fileName)
	expected := 0

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}

// test for shredding a file which does not have a write permission
func Test_3(t *testing.T) {
	fileName := "../test_files/test_file_2.txt"

	result := shred(fileName)
	expected := 3

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}