package main

import (
	"os"
	"fmt"
	"os/exec"
	"io/ioutil"
)

func SaveReadFile(data []byte)  {
	var file *os.File
	file, err := os.Create("voice.m4a")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(data))

	fmt.Println("File Save Exec")
}

func ReadWav() []byte{
	data, err := ioutil.ReadFile(`./voice.wav`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

// TODO 未実装
func ConvertM4aToWav() (error) {
	err := exec.Command("ffmpeg -i voice.m4a voice.wav").Run()
	if err != nil {
		return err
	}
	return nil
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