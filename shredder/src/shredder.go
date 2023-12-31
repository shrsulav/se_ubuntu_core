package shredder

import (
	"os"
	"math/rand"
	"log"
	"path/filepath"
)

const maxShredCount uint 	= 3
const maxShredBytes int 	= 134_217_728	// 128MB

// function to check if the parent directory for the given file has exec permissions
func HasExecPerm(fileName string) *ShredderError {

	// not checking if the file exists in this function
	// has to be checked before the function is called

	dir, _ := filepath.Split(fileName)

	cwDir, wdErr := os.Getwd()
	if wdErr != nil {
		log.Println(wdErr)
		err := ReturnInfo(ShredErrProcessing, ShredErrProcessing.ShredErrString())
		return err
	}

	dirError := os.Chdir(dir)
	if dirError != nil {
		log.Printf("%v\n", dirError)
		err := ReturnInfo(ShredErrNoExecutePerm, ShredErrNoExecutePerm.ShredErrString())
		return err
	}
	_ = os.Chdir(cwDir)

	return nil
}

func WriteToFileHandle(fileHandle *os.File, fileBytesNum int64, randomDataSize int) error {

	var shreddedBytes int64 = 0

	randomData := make([]byte, randomDataSize)

	for shreddedBytes < fileBytesNum {
		// generate random data
		_, randErr := rand.Read(randomData)

		if randErr != nil {
			log.Printf("Error generating random data.\n");
		}

		// write to file
		numBytes, writeError := fileHandle.WriteAt(randomData, int64(shreddedBytes))

		if writeError != nil {
			return writeError
		}

		shreddedBytes += int64(numBytes)
	}

	return nil
}

func Shred(fileName string) *ShredderError {
	log.Printf("Shredding file: %v\n", fileName)

	// check if the file exists
	fileInfo, fileError := os.Stat(fileName)

	if os.IsNotExist(fileError) {
		log.Printf("\"%v\" file does not exist.\n", fileName)
		err := ReturnInfo(ShredErrFileNotExist, ShredErrFileNotExist.ShredErrString())
		return err
	}

	// check if the parent directory has executable permission
	permError := HasExecPerm(fileName)
	if permError != nil	{
		return permError
	}

	// check if the given path is a file or a directory
	if !fileInfo.Mode().IsRegular() {
        log.Println(fileName, "is not a regular file.")
		err := ReturnInfo(ShredErrNotAFile, ShredErrNotAFile.ShredErrString())
		return err
    }

	fileBytesNum := fileInfo.Size()
	log.Printf("The file \"%v\" is %d bytes long.\n", fileName, fileBytesNum)

	var randomDataSize int

	if fileBytesNum <= 0 {
		log.Println("Error: file size is zero or negative.")
		randomDataSize = rand.Int()
	} else if fileBytesNum > int64(maxShredBytes) {
		log.Println("File size is greater than maximum shred bytes (128MB).")
		randomDataSize = maxShredBytes
	} else {
		randomDataSize = int(fileBytesNum)
	}

	log.Printf("Random data size is %d\n", randomDataSize)

	shredCount := maxShredCount
	shredErrCount := 0

	for shredCount > 0 {
		fileHandle, openError := os.OpenFile(fileName, os.O_RDWR, 0666)

		if openError != nil {
			log.Printf("%v\n", openError)
			err := ReturnInfo(ShredErrFileOpen, ShredErrFileOpen.ShredErrString())
			return err
		}

		defer fileHandle.Close()

		randomizeErr := WriteToFileHandle(fileHandle, fileBytesNum, randomDataSize)
		if randomizeErr != nil {
			shredErrCount += 1
		}
		shredCount -= 1
	}

	if shredErrCount != 0 {
		log.Printf("Error shredding!\n")
		err := ReturnInfo(ShredErrFileWrite, ShredErrFileWrite.ShredErrString())
		return err
	}

	// delete the file
	delErr := os.Remove(fileName)
    if delErr != nil {
        log.Printf("%v\n", delErr)
		err := ReturnInfo(ShredErrFileDelete, ShredErrFileDelete.ShredErrString())
		return err
    } else {
		log.Printf("Deleted the file \"%v\"\n", fileName)
	}

	err := ReturnInfo(ShredErrSuccess, ShredErrSuccess.ShredErrString())
	return err
}


