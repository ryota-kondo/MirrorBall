package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"bytes"
	"path/filepath"
	"io"
	"mime/multipart"
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
	var err error

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("wav", filepath.Base("./voice.wav"))
	if err != nil {
		return EmpathResponse{}
	}
	_, err = io.Copy(part, bytes.NewReader(voice))


	_ = writer.WriteField("apikey", apiKey)

	err = writer.Close()
	if err != nil {
		return EmpathResponse{}
	}

	req, err := http.NewRequest("POST", "https://api.webempath.net/v2/analyzeWav", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return EmpathResponse{}
	} 
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
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
	fmt.Println("api ok")
	return empathResponse

	//body := bytes.Buffer{}
	//writer := multipart.NewWriter(&body)
	//part1, err := writer.CreateFormFile("wav", filepath.Base("voice.wav"))
	//if err != nil {
	//	fmt.Println("api 1-1 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}
	//_, err =io.Copy(part1, bytes.NewReader(voice))
	//if err != nil {
	//	fmt.Println("api 1-2 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}

	//part2, err := writer.CreateFormField("apiKey")
	//if err != nil {
	//	fmt.Println("api 1-4 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}
	//_, err =io.Copy(part2, strings.NewReader(apiKey))
	//if err != nil {
	//	fmt.Println("api 1-5 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}

	//writer.WriteField("apiKey",apiKey)
	//
	//err = writer.Close()
	//if err != nil {
	//	fmt.Println("api 1-6 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}
	//
	//fmt.Println(fmt.Sprintf("%+v",writer))

	//r, _ := http.NewRequest("POST", "https://api.webempath.net/v2/analyzeWav", &body)
	//r.Header.Add("Content-Type", writer.FormDataContentType())
	//client := &http.Client{}
	//res, err := client.Do(r)
	//if err != nil {
	//	fmt.Println("api 1 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}
	//defer res.Body.Close()
	//
	//resBody, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("api 2 error")
	//	fmt.Println(err)
	//	return EmpathResponse{}
	//}
	//
	//var empathResponse EmpathResponse
	//err = json.Unmarshal(resBody, &empathResponse)
	//if err != nil {
	//	fmt.Println("api 3 error")
	//	fmt.Println(err)
	//	fmt.Println(string(resBody))
	//	return EmpathResponse{}
	//}
	//
	//fmt.Println("JSON")
	//fmt.Println(fmt.Sprintf("%v",empathResponse))
	//
	////empathResponse.Error = 0
	////empathResponse.Message = "test"
	////empathResponse.Anger = 25
	////empathResponse.Calm = 50
	////empathResponse.Energy = 9
	////empathResponse.Joy = 42
	////empathResponse.Sorrow = 5
	//
	//return empathResponse
}