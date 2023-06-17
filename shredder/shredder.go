package shredder

import "os"
import "math/rand"
import "log"

func shred(fileName string) int {
	log.Printf("Shredding file: %v\n", fileName)

	// check if the file exists

	_, fileError := os.Stat(fileName)

	if os.IsNotExist(fileError) {
		log.Printf("\"%v\" file does not exist.\n", fileName)
		return -1
	} else {
		log.Printf("\"%v\" file exists.\n", fileName)
	}

	// log.Printf("The file \"%v\" is %d bytes long.\n", fileName, fileInfo.Size())

	// var randomDataSize uint64

	// if(fileInfo.Size() <= 0) {
	// 	log.Println("Error: file size is zero or negative.")
	// 	randomDataSize = rand.Uint64()
	// } else {
	// 	randomDataSize = uint64(fileInfo.Size())
	// }

	// log.Printf("Random data size is %d\n", randomDataSize)

	randomData := make([]byte, 1000)

	shredCount := 3
	shredErrCount := 0

	for shredCount > 0 {

		// generate random data

		_, err := rand.Read(randomData)

		if err != nil {
			log.Printf("Error generating random data.\n");
		}

		// write to file

		writeError := os.WriteFile(fileName, randomData, 666)

		if writeError != nil {
			log.Printf("Shredder Pass %d : %v\n", 4 - shredCount, writeError)
			shredErrCount += 1
		} else {
			log.Printf("Shredding successful.\n")
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


