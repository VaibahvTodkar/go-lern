package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File in the GO language")
	content := "This is my first Go program."

	file, err := os.Create("./myfile.txt")
	checkError(err)
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	panic(err) // print error
	// 	return
	// }

	length, err := io.WriteString(file, content)
	checkError(err)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	panic(err)
	// 	return
	// }

	fmt.Printf("Wrote %d bytes to file.\n", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkError(err)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	panic(err)
	// 	return
	// }
	fmt.Println("File content:", string(dataByte))
	// dataByte   [84 104 105 115 32 105 115 32 109 121 32 102 105 114 115 116 32 71 111 32 112 114 111 103 114 97 109 46]
	// string(dataByte) // This is my first Go program.
}

// use single code to check error handling
func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
