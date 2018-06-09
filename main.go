package main

import (
	"io"
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"bytes"
)

var codec = websocket.Codec{}

func main() {
	http.Handle("/mirror_bowl", websocket.Handler(MirrorBowlHandler))
	http.Handle("/echo", websocket.Handler(EchoHandler))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func MirrorBowlHandler(ws *websocket.Conn) {
	//str,_ := ioutil.ReadAll(ws)
	//buffer3 := bytes.NewBufferString(string(str))
	buffer3 := bytes.NewBufferString("a")
	ws.Write(buffer3.Bytes())
}

func EchoHandler(ws *websocket.Conn) {
	fmt.Println("test")
	io.Copy(ws, ws)
}
