package main

import (
	"os"
	"log"
	"fmt"
)

func SaveReadFile(data []byte)  {
	var file *os.File
	file, err := os.Create("tmp")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(data))

	fmt.Println("File Save Exec")
}

func ReadWav() *os.File{
	// ファイルのオープ
	file, err := os.Open("./tmp")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func ConvertM4aToWav()  {
	
}