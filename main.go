package main

import (
	"os"
	"io"
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"encoding/base64"
)

// 前回沈黙フラグ
var latestResult = false

//　クライアントから送信されるエンコード済み音声データ
type Voice struct {
	Data string`json:"data"`
}

type Emotion struct {
	Calm   int `json:"calm"`
	Anger  int `json:"anger"`
	Joy    int `json:"joy"`
	Sorrow int `json:"sorrow"`
	Energy int `json:"energy"`
}

type MirrorBallResponse struct {
	Suggestion string  `json:"suggestion"`
	Tention    int     `json:"tention"`
	Emotion    Emotion `json:"emotion"`
}

func main() {
	fmt.Println("Run Server")

	http.Handle("/mirror_ball", websocket.Handler(MirrorBallHandler))
	http.Handle("/echo", websocket.Handler(EchoHandler))
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func MirrorBallHandler(ws *websocket.Conn) {
	var err error
	for {
		var v Voice
		if err = websocket.JSON.Receive(ws, &v); err != nil {
			fmt.Println("[E0001]")
			fmt.Println(err)
			break
		}

		// 音声をBase64でコード
		data, _ := base64.StdEncoding.DecodeString(v.Data) //[]byte

		// m4aファイルをローカルへ保存
		SaveReadFile(data)

		// m4a -> WAVのシェルを実行(FFMPEG)
		ConvertM4aToWav()

		// 変換したwavファイル読み込み
		wavData := ReadWav()

		// EmpathAPIにWAVを送信し感情を受け取る
		empath := SendEmpathAPI(wavData)

		// API結果よりResponseJパラメータを生成
		res := CreateResponse(empath)

		// クライアントに送信
		if err = websocket.JSON.Send(ws, res); err != nil {
			fmt.Println("[E0002]")
			fmt.Println(err)
		}
	}
}

// おうむ返し
func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
