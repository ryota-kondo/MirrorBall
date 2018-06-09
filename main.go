package main

import (
	"io"
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"io/ioutil"
)

type Voice struct {
	Data []byte
}

type Emotion struct{
	Calm    int    `json:"calm"`
	Anger   int    `json:"anger"`
	Joy     int    `json:"joy"`
	Sorrow  int    `json:"sorrow"`
	Energy  int    `json:"energy"`
}

type MirrorBowlResponse struct {
	Suggestion string `json:"suggestion"`
	Tention    int    `json:"tention"`
	Emotion Emotion`json:"emotion"`
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
		//var v Voice
		//if err = websocket.JSON.Receive(ws, &v); err != nil {
		//	fmt.Println(err)
		//	break
		//}
		var v string
		if err = websocket.Message.Receive(ws, &v); err != nil {
			fmt.Println(err)
			break
		}

		data, err := ioutil.ReadFile(`./voice.wav`)
		if err != nil {
			fmt.Println(err)
			break
		}

		empath := SendEmpathAPI(data)

		res := CreateResponse(empath)

		//data = T{
		//	Msg:   message,
		//	Count: 114514,
		//}

		if err = websocket.JSON.Send(ws, res); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
