package shredder

import "fmt"
import "os"
import "math/rand"

func shred(fileName string) int {
	fmt.Printf("Shredding file: %v\n", fileName)

	// check if the file exists

	fileInfo, fileError := os.Stat(fileName)

	if os.IsNotExist(fileError) {
		fmt.Printf("\"%v\" file does not exist.\n", fileName)
		return -1
	} else {
		fmt.Printf("\"%v\" file exists.\n", fileName)
	}

	fmt.Printf("The file \"%v\" is %d bytes long.\n", fileName, fileInfo.Size())

	var randomDataSize uint64

	if(fileInfo.Size() <= 0) {
		fmt.Println("Error: file size is zero or negative.")
		randomDataSize = rand.Uint64()
	} else {
		randomDataSize = uint64(fileInfo.Size())
	}

	fmt.Printf("Random data size is %d\n", randomDataSize)

	randomData := make([]byte, 1000)

	shredCount := 3
	shredErrCount := 0

	for shredCount > 0 {

		// generate random data

		_, err := rand.Read(randomData)

		if err != nil {
			fmt.Printf("Error generating random data.\n");
		}

		// write to file

		writeError := os.WriteFile(fileName, randomData, 666)

		if writeError != nil {
			fmt.Printf("Shredder Pass %d : %v\n", 4 - shredCount, writeError)
			shredErrCount += 1
		} else {
			fmt.Printf("Shredding successful.\n")
		}

		shredCount -= 1
	}

	// delete the file
	delErr := os.Remove(fileName)
    if delErr != nil {
        fmt.Printf("%v\n", delErr)
    } else {
		fmt.Printf("Deleted the file \"%v\"\n", fileName)
	}

	return shredErrCount
}


