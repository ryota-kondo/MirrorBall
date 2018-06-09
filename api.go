package main

import (
	"mime/multipart"
	"path/filepath"
	"io"
	"net/http"
	"bytes"
	"strings"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type EmpathRequest struct {
	ApiKey string     `json:"apiKey"`
	Wav    []byte `json:"wav"`
}

type EmpathResponse struct {
	Error   int `json:"error"`
	Message string `json:"msg"`
	Calm    int    `json:"calm"`
	Anger   int    `json:"anger"`
	Joy     int    `json:"joy"`
	Sorrow  int    `json:"sorrow"`
	Energy  int    `json:"energy"`
}

func SendEmpathAPI(voice []byte) EmpathResponse{
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part1, _ := writer.CreateFormFile("wav", filepath.Base("voice.wav"))
	io.Copy(part1, bytes.NewReader(voice))
	part2, _ := writer.CreateFormField("apiKey")
	io.Copy(part2, strings.NewReader(apiKey))

	writer.Close()

	r, _ := http.NewRequest("POST", "https://api.webempath.net/v2/analyzeWav", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("api 1 error")
		fmt.Println(err)
		return EmpathResponse{}
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("api 2 error")
		fmt.Println(err)
		return EmpathResponse{}
	}

	var empathResponse EmpathResponse
	err = json.Unmarshal(resBody, &empathResponse)
	if err != nil {
		fmt.Println("api 3 error")
		fmt.Println(err)
		fmt.Println(string(resBody))
		return EmpathResponse{}
	}

	fmt.Println("JSON")
	fmt.Println(fmt.Sprintf("%v",empathResponse))

	//empathResponse.Error = 0
	//empathResponse.Message = "test"
	//empathResponse.Anger = 25
	//empathResponse.Calm = 50
	//empathResponse.Energy = 9
	//empathResponse.Joy = 42
	//empathResponse.Sorrow = 5

	return empathResponse
}