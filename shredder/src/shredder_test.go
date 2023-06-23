package shredder

import (
	"testing"
	"os"
	"log"
	"math/rand"
)

// test for shredding a file which does not exist
func Test_1(t *testing.T) {
	createErr := os.MkdirAll("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
		return
	}

	fileName := "testDir/nonexistentfile.txt"

	result := Shred(fileName)
	expected := ShredErrFileNotExist

	if result.ErrCode != expected {
		t.Errorf("got %v, expected %v", result.ErrCode, expected)
	}

	removeErr := os.Remove("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}

// test for shredding a file which exists and has rw file permissions
func Test_2(t *testing.T) {
	createErr := os.MkdirAll("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
		return
	}

	fileName := "testDir/test_file_1.txt"

	randomData := make([]byte, 100)

	_, randError := rand.Read(randomData)
	if randError != nil {
		log.Printf("Error while generating random string: %s", randError)
	}

	writeError := os.WriteFile(fileName, randomData, 0666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		} else {
			log.Printf("Successfully deleted test directory.")
		}

		return
	}
	result := Shred(fileName)
	expected := ShredErrSuccess

	if result.ErrCode != expected {
		t.Errorf("got %+v, expected %+v", result.ErrCode, expected)
	}

	// check if the file exists
	_, fileError := os.Stat(fileName)

	if !os.IsNotExist(fileError) {
		t.Errorf("File should have been deleted but is not.")
	} else {
		log.Println("File deleted.")
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}

// test for shredding a file which does not have a write permission
func Test_3(t *testing.T) {
	createErr := os.MkdirAll("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
	}

	fileName := "testDir/test_file_2.txt"

	randomData := make([]byte, 100)

	_, randError := rand.Read(randomData)
	if randError != nil {
		log.Printf("Error while generating random string: %s", randError)
	}

	writeError := os.WriteFile(fileName, randomData, 0666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		} else {
			log.Printf("Successfully deleted test directory.")
		}

		return
	}

	modErr := os.Chmod(fileName, 0444)
   	if modErr != nil {
      	t.Errorf("Error making file read-only")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		} else {
			log.Printf("Successfully deleted test directory.")
		}

      	return
   	}

	result := Shred(fileName)
	expected := ShredErrFileOpen

	if result.ErrCode != expected {
		t.Errorf("got %+v, expected %+v", result.ErrCode, expected)
	}

	// check if the file exists
	_, fileError := os.Stat(fileName)
	if os.IsNotExist(fileError) {
		t.Errorf("File should not have been deleted, but is.")
	} else {
		log.Println("File not deleted.")
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}

// test for shredding a file owned by the root
// func Test_4(t *testing.T) {

// 	createErr := os.MkdirAll("testDir", 0777)

// 	if createErr != nil {
// 		t.Errorf("Error creating test directory.")
// 	}

// 	fileName := "testDir/test_file_3.txt"

// 	randomData := make([]byte, 100)

// 	_, randError := rand.Read(randomData)
// 	if randError != nil {
// 		log.Printf("Error while generating random string: %s", randError)
// 	}

// 	writeError := os.WriteFile(fileName, randomData, 0666)

// 	if writeError != nil {
// 		t.Errorf("Error creating test file.")

// 		removeErr := os.RemoveAll("testDir")

// 		if removeErr != nil {
// 			t.Errorf("Error deleting the test directory")
// 		} else {
// 			log.Printf("Successfully deleted test directory.")
// 		}

// 		return
// 	}

// 	// change file ownership to root
// 	ownErr := os.Chown(fileName, root, root )
// 	if ownErr != nil {
// 		t.Errorf("Error changing file ownership")

// 		removeErr := os.RemoveAll("testDir")

// 		if removeErr != nil {
// 			t.Errorf("Error deleting the test directory")
// 		} else {
// 			log.Printf("Successfully deleted test directory.")
// 		}

// 		return
// 	}

// 	result := Shred(fileName)
// 	expected := ShredErrSuccess

// 	if result.ErrCode != expected {
// 		t.Errorf("got %+v, expected %+v", result.ErrCode, expected)
// 	}

// 	removeErr := os.RemoveAll("testDir")

// 	if removeErr != nil {
// 		t.Errorf("Error deleting the test directory")
// 	} else {
// 		log.Printf("Successfully deleted test directory.")
// 	}
// }

// test for shredding a file in a directory which does not have executable permission
func Test_5(t *testing.T) {

	createErr := os.MkdirAll("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
	}

	fileName := "testDir/test_file_4.txt"

	randomData := make([]byte, 100)

	_, randError := rand.Read(randomData)
	if randError != nil {
		log.Printf("Error while generating random string: %s", randError)
	}

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
	result := Shred(fileName)
	expected := ShredErrNoExecutePerm

	if result.ErrCode != expected {
		t.Errorf("got %+v, expected %+v", result.ErrCode, expected)
	}

	remodError := os.Chmod("testDir", 0777)

	if remodError != nil {
		t.Errorf("Error readding executable permissions from the test directory.")
		log.Printf("%v\n", remodError)
	}

	// check if the file exists
	_, fileError := os.Stat(fileName)
	if os.IsNotExist(fileError) {
		t.Errorf("File should not have been deleted but is.")
	} else {
		log.Println("File not deleted.")
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}

// test for shredding a directory instead of a file
func Test_6(t *testing.T) {

	createErr := os.MkdirAll("testDir/testSubDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
	}

	fileName := "testDir/testSubDir/test_file_4.txt"

	randomData := make([]byte, 100)

	writeError := os.WriteFile(fileName, randomData, 0666)

	_, randError := rand.Read(randomData)
	if randError != nil {
		log.Printf("Error while generating random string: %s", randError)
	}

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		}

		return
	}

	result := Shred("testDir/testSubDir")
	expected := ShredErrNotAFile

	if result.ErrCode != expected {
		t.Errorf("got %+v, expected %+v", result.ErrCode, expected)
	}

	// check if the file exists
	_, fileError := os.Stat(fileName)
	if os.IsNotExist(fileError) {
		t.Errorf("File should not have been deleted but is.")
	} else {
		log.Println("File not deleted.")
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory.")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}

// test for shredding a file which is larger than 128MB
func Test_7(t *testing.T) {
	createErr := os.MkdirAll("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
		return
	}

	fileName := "testDir/test_file_7.txt"

	randomData := make([]byte, 314_572_800) // 300MB

	_, randError := rand.Read(randomData)
	if randError != nil {
		log.Printf("Error while generating random string: %s", randError)
	}

	writeError := os.WriteFile(fileName, randomData, 0666)

	if writeError != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		} else {
			log.Printf("Successfully deleted test directory.")
		}

		return
	}
	result := Shred(fileName)
	expected := ShredErrSuccess

	if result.ErrCode != expected {
		t.Errorf("got %+v, expected %+v", result.ErrCode, expected)
	}

	// check if the file exists
	_, fileError := os.Stat(fileName)
	if !os.IsNotExist(fileError) {
		t.Errorf("File should have been deleted but is not.")
	} else {
		log.Println("File deleted.")
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}

// test if the helper function actually writes random data to the given file handle
func Test_8(t *testing.T) {
	createErr := os.MkdirAll("testDir", 0777)

	if createErr != nil {
		t.Errorf("Error creating test directory.")
		return
	}

	fileName := "testDir/test_file_8.txt"
	fileNameBkp := "testDir/test_file_8_bkp.txt"

	fileSize := 10_000
	randomData := make([]byte, fileSize)

	_, randError := rand.Read(randomData)
	if randError != nil {
		log.Printf("Error while generating random string: %s", randError)
	}

	writeError := os.WriteFile(fileName, randomData, 0666)
	writeError2 := os.WriteFile(fileNameBkp, randomData, 0666)

	if writeError != nil && writeError2 != nil {
		t.Errorf("Error creating test file.")

		removeErr := os.RemoveAll("testDir")

		if removeErr != nil {
			t.Errorf("Error deleting the test directory")
		} else {
			log.Printf("Successfully deleted test directory.")
		}

		return
	}

	fileHandle, openError := os.OpenFile(fileName, os.O_RDWR, 0666)

	if openError != nil {
		log.Printf("%v\n", openError)
		t.Errorf("Error while opening file.\n")
	}

	defer fileHandle.Close()

	result := WriteToFileHandle(fileHandle, int64(fileSize), 100)

	if result != nil {
		t.Errorf("WriteToFileHandle not successful")
	}

	dataRandomized, randomizedFileErr := os.ReadFile(fileName)
	dataOriginal, originalFileErr := os.ReadFile(fileNameBkp)

	if randomizedFileErr != nil && originalFileErr != nil {
		t.Errorf("Error while checking content randomization.")
	}
	log.Println("Number of bytes in original file: ", len(dataOriginal))
	log.Println("Number of bytes in randomized file: ", len(dataRandomized))

	if len(dataOriginal) > len(dataRandomized) {
		t.Errorf("Less data bytes after randomization")
	}

	var diffCount int = 0

	for i:=0; i<len(dataOriginal); i++ {
		if dataOriginal[i] != dataRandomized[i] {
			diffCount++
		}
	}

	log.Println("Diff count: ", diffCount)

	randomizationPercent := diffCount * 100 / fileSize
	log.Println("Randomization percentage is ", randomizationPercent)
	if randomizationPercent < 75  {
		t.Errorf("Randomization is less than 75")
	}

	removeErr := os.RemoveAll("testDir")

	if removeErr != nil {
		t.Errorf("Error deleting the test directory")
	} else {
		log.Printf("Successfully deleted test directory.")
	}
}
