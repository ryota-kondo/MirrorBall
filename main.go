package main

import (
	"io"
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
)

type T struct {
	Msg string
	Count int
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
		var message string
		if err = websocket.Message.Receive(ws, &message); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + message)

		data := T{
			Msg:   message,
			Count: 1,
		}

		if err = websocket.JSON.Send(ws, data); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
