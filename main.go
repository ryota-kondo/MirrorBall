package main

import (
	"io"
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"encoding/base64"
)

// 前回沈黙フラグ
var latestResult = false

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

type MirrorBowlResponse struct {
	Suggestion string  `json:"suggestion"`
	Tention    int     `json:"tention"`
	Emotion    Emotion `json:"emotion"`
}

func main() {
	fmt.Println("Run Server")

	http.Handle("/mirror_bowl", websocket.Handler(MirrorBowlHandler))
	http.Handle("/echo", websocket.Handler(EchoHandler))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func MirrorBowlHandler(ws *websocket.Conn) {
	var err error
	for {
		var v Voice
		if err = websocket.JSON.Receive(ws, &v); err != nil {
			fmt.Println("[E0001]")
			fmt.Println(err)
			break
		}

		data, _ := base64.StdEncoding.DecodeString(v.Data) //[]byte

		// m4aファイルを保存
		SaveReadFile(data)

		// m4a -> WAV
		ConvertM4aToWav()

		// wavファイル読み込み
		wavData := ReadWav()

		// API 叩き
		empath := SendEmpathAPI(wavData)

		res := CreateResponse(empath)

		if err = websocket.JSON.Send(ws, res); err != nil {
			fmt.Println("[E0002]")
			fmt.Println(err)
		}
	}
}

func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
