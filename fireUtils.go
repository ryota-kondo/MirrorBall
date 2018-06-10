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
	fmt.Println("cnv start")

	out, err := exec.Command("ls", "-la").Output()
	fmt.Println(string(out))

	err = exec.Command("sh","./enc.sh").Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("conv Success")
	return nil
}

func CreateResponse(empathResponse EmpathResponse,pt int) MirrorBowlResponse {
	var res MirrorBowlResponse

	if pt == 0 {
		res = MirrorBowlResponse{
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
	}else if pt == 1{
		res = MirrorBowlResponse{
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
	}else if pt == 2{

	}



	return res
}