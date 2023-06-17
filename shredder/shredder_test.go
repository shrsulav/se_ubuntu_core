package shredder

import "testing"
import "os"

// test for shredding a file which does not exist
func Test_1(t *testing.T) {
	createErr := os.Mkdir("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
		return
	}

	fileName := "testDir/nonexistentfile.txt"

	result := shred(fileName)
	expected := -1

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}

	removeErr := os.Remove("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	}
}

// test for shredding a file which exists and has rw file permissions
func Test_2(t *testing.T) {
	createErr := os.Mkdir("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
		return
	}

	fileName := "testDir/test_file_1.txt"

	randomData := make([]byte, 1000)

	writeError := os.WriteFile(fileName, randomData, 666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}
	result := shred(fileName)
	expected := 0

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	}
}

// test for shredding a file which does not have a write permission
func Test_3(t *testing.T) {
	createErr := os.Mkdir("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
	}

	fileName := "testDir/test_file_2.txt"

	randomData := make([]byte, 1000)

	writeError := os.WriteFile(fileName, randomData, 0666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}

	modErr := os.Chmod(fileName, 0444)
   	if modErr != nil {
      	t.Errorf("Error making file read-only")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

      	return
   	}

	result := shred(fileName)
	expected := 3

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	}
}

// test for shredding a file owned by the root
func Test_4(t *testing.T) {

	createErr := os.Mkdir("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
	}

	fileName := "testDir/test_file_3.txt"

	randomData := make([]byte, 1000)

	writeError := os.WriteFile(fileName, randomData, 0666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}

	// change file ownership to root
	ownErr := os.Chown(fileName, 1000, 1000 )
	if ownErr != nil {
		t.Errorf("Error changing file ownership")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}

	result := shred(fileName)
	expected := 0

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	}
}

// test for shredding a file in a directory which does not have executable permission
func Test_5(t *testing.T) {

	createErr := os.Mkdir("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
	}

	fileName := "testDir/test_file_4.txt"

	randomData := make([]byte, 1000)

	writeError := os.WriteFile(fileName, randomData, 0666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}

	modError := os.Chmod("testDir", 0666)

	if modError != nil {
		t.Errorf("Error removing executable permissions from the test directory.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}
	result := shred(fileName)
	expected := 3

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}

	remodError := os.Chmod("testDir", 0777)

	if remodError != nil {
		t.Errorf("Error removing executable permissions from the test directory.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	}
}