// execute "go run shredder.go tester.go" to execute the program

package main

import "fmt"

func main() {
  fmt.Println("Hello World!")
  result := shred("textfile.txt")
  fmt.Println("Shredding result: ", result)
}

