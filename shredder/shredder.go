package shredder

import (
	"os"
	"math/rand"
	"log"
	"path/filepath"
)

func Shred(fileName string) int {
	log.Printf("Shredding file: %v\n", fileName)

	// check if the file exists

	fileInfo, fileError := os.Stat(fileName)

	if os.IsNotExist(fileError) {
		log.Printf("\"%v\" file does not exist.\n", fileName)
		return -1
	} else {
		log.Printf("\"%v\" file exists.\n", fileName)
	}

	dir, _ := filepath.Split(fileName)
	log.Println("Directory name is :", dir)

	dirError := os.Chdir(dir)
	if dirError != nil {
		log.Printf("%v\n", dirError)
		return -1
	}
	_ = os.Chdir("..")

	if !fileInfo.Mode().IsRegular() {
        log.Println(fileName, "is not a regular file!")
		return -1
    }

	log.Printf("The file \"%v\" is %d bytes long.\n", fileName, fileInfo.Size())

	var randomDataSize int

	if(fileInfo.Size() <= 0) {
		log.Println("Error: file size is zero or negative.")
		randomDataSize = rand.Int()
	} else {
		randomDataSize = int(fileInfo.Size())
	}

	log.Printf("Random data size is %d\n", randomDataSize)

	randomData := make([]byte, 1000)

	shredCount := 3
	shredErrCount := 0

	fileHandle, openError := os.OpenFile(fileName, os.O_RDWR, 0666)

	if openError != nil {
		log.Printf("%v\n", openError)
		return -1
	}

	defer fileHandle.Close()

	var readBuffer []byte
	_, readError := fileHandle.Read(readBuffer)

	if readError != nil {
		log.Printf("%v\n", readError)
		return -1
	}

	readStringBuffer := string(readBuffer)
	log.Println(readStringBuffer)

	for shredCount > 0 {

		var shreddedBytes int = 0

		for shreddedBytes < randomDataSize {

			// generate random data
			_, randErr := rand.Read(randomData)

			if randErr != nil {
				log.Printf("Error generating random data.\n");
			}

			// write to file
			numBytes, writeError := fileHandle.WriteAt(randomData, int64(shreddedBytes))

			if writeError != nil {
				log.Printf("Shredder Pass %d : %v\n", 4 - shredCount, writeError)
				shredErrCount += 1
			} else {
				log.Printf("Shredder Pass %d: Shredding successful.\n", 4 - shredCount)
			}

			shreddedBytes += numBytes

		}

		shredCount -= 1
	}

	if shredErrCount != 0 {
		log.Printf("Error shredding!\n")
		return shredErrCount
	}

	// delete the file

	delErr := os.Remove(fileName)
    if delErr != nil {
        log.Printf("%v\n", delErr)
		return 4
    } else {
		log.Printf("Deleted the file \"%v\"\n", fileName)
	}

	return 0
}


