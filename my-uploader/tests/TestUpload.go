package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"bytes"
	"mime/multipart"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	_, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	//fh, err := os.Open(filename)
	//if err != nil {
	//	fmt.Println("error opening file")
	//	return err
	//}
	//defer fh.Close()
	//
	////iocopy
	//_, err = io.Copy(fileWriter, fh)
	//if err != nil {
	//	return err
	//}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	println(contentType)

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}


func httpPost() {
	filename := "/Users/adhoc-dev/Documents/adhoc-perf-test/Result/6000-gateway1-exp1-tracker1-resp.png"
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
	}

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/upload/0e13d87e-fa04-4a07-bdd5-65512f870ea5", strings.NewReader("name=cjb"))

	if err != nil {
		fmt.Println(err)
	}
	println(contentType)
	request.Header.Set("Auth-Key", "49ec7d89-3210-4cd6-bdb4-bde410be3cb1")
	request.Header.Add("Content-Type", contentType)

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	filename := "/Users/adhoc-dev/Documents/adhoc-arch-graph.jpg"
	url := "http://localhost:8080/upload/0e13d87e-fa04-4a07-bdd5-65512f870ea5/"
	//url := "http://localhost:8080/upload"
	postFile(filename, url)

	//httpPost()

}

