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
	file, err := os.Open("./tmp")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// TODO 未実装
func ConvertM4aToWav()  {





}

func CreateResponse(empathResponse EmpathResponse) MirrorBowlResponse {
	return MirrorBowlResponse{
		Suggestion:"森食ってモリモリ",
		Tention:49,
		Emotion: Emotion{
			empathResponse.Calm ,
			empathResponse.Anger ,
			empathResponse.Joy ,
			empathResponse.Sorrow ,
			empathResponse.Energy,
		},
	}
}