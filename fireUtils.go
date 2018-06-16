package main

import (
	"os"
	"fmt"
	"os/exec"
	"io/ioutil"
)

var Tension [2]int
var ResponseCount int

// m4a音声をローカルに保存
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

// WAVをローカルから読み取り
func ReadWav() []byte{
	data, err := ioutil.ReadFile(`./voice.wav`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

// m4aをwavに変換するシェルスクリプトを実行
func ConvertM4aToWav() (error) {
	fmt.Println("cnv start")

	err := exec.Command("sh","./enc.sh").Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("conv Success")
	return nil
}


func CreateResponse(empathResponse EmpathResponse) MirrorBallResponse {

	// 統計処理
	Tension[0] = Tension[1]
	Tension[1] = empathResponse.Energy
	ResponseCount += 1

	if ResponseCount > 1 && ResponseCount % 2 == 0 {
		var resTension = (Tension[0] + Tension[1])/2

		// 評価回数である場合

		// APIエラーがある場合
		if empathResponse.Error != 0 {
			return MirrorBallResponse{
				Suggestion: "9",
				Tention:0,
				Emotion: Emotion{
					0,
					0,
					0,
					0,
					0,
				},
			}
			// APIエラーがない場合
		} else {
			// init
			var checkArry = []int{empathResponse.Calm, empathResponse.Anger, empathResponse.Joy, empathResponse.Sorrow }
			var checkSuggestion [2]string
			var resSuggestion string

			// Max感情を判定
			// Max感情を判定
			var res int
			for i , v := range checkArry {
				var checkValue int
				if checkValue < v {
					res = i
					checkValue = v
				}
			}

			// 感情分類
			switch res {
			case 0:
				checkSuggestion[0] = "1"
				checkSuggestion[1] =  "5"
			case 1:
				checkSuggestion[0] = "2"
				checkSuggestion[1] = "6"
			case 2:
				checkSuggestion[0] = "3"
				checkSuggestion[1] = "7"
			case 3:
				checkSuggestion[0] = "4"
				checkSuggestion[1] = "8"
			}

			switch {
			case resTension < 20:
				resSuggestion = checkSuggestion[1]
			case resTension >= 20:
				resSuggestion = checkSuggestion[0]
			}

			return MirrorBallResponse{
				Suggestion: resSuggestion,
				Tention: resTension,
				Emotion: Emotion{
					empathResponse.Calm ,
					empathResponse.Anger ,
					empathResponse.Joy ,
					empathResponse.Sorrow ,
					empathResponse.Energy,
				},
			}
		}
	} else {
		// 評価回数じゃない場合
		return MirrorBallResponse{
			Suggestion: "",
			Tention: empathResponse.Energy,
			Emotion: Emotion{
				empathResponse.Calm ,
				empathResponse.Anger ,
				empathResponse.Joy ,
				empathResponse.Sorrow ,
				empathResponse.Energy,
			},
		}
	}
}