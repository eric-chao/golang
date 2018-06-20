package handler

import (
	router "github.com/julienschmidt/httprouter"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	. "adhoc/adhoc_uploader/config"
	. "adhoc/adhoc_uploader/logging"
	. "adhoc/adhoc_uploader/resp"
)

var AcceptContentType = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
}

func HttpOption(w http.ResponseWriter, r *http.Request, p router.Params) {
	//sends data to client side
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Auth-Key")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}

func SayHello(w http.ResponseWriter, r *http.Request, p router.Params) {
	//sends data to client side
	fmt.Fprintf(w, "Hello Adhoc!")
}

func Configuration(w http.ResponseWriter, r *http.Request, p router.Params) {
	result, _ := json.Marshal(GlobalConfig)
	w.Write(result)
}

func DoUpload(w http.ResponseWriter, r *http.Request, p router.Params) {
	//max memory on server - 2MB
	r.ParseMultipartForm(2 << 20)
	// get file handle
	file, fileHeader, err := r.FormFile("uploadfile")
	if err != nil {
		Logger.Error(err.Error())
		ResponseBody(w, 90000, "read upload-file error!", "")
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err == nil {
		// max file size: 300k
		if len(data) > 300<<10 {
			ResponseBody(w, 91000, "maxFileSize more than 300k.", "")
			return
		}
	}

	Logger.Infof("%v", fileHeader.Header)
	if _, ok := AcceptContentType[fileHeader.Header.Get("Content-Type")]; !ok {
		ResponseBody(w, 91000, "only [.jpg .png .gif] supported.", "")
		return
	}

	appId := p.ByName("appId")
	// make sure GlobalConfig.Storage.Path ended with '/'
	outputPath := fmt.Sprintf("%s%s/", GlobalConfig.Storage.Path, appId)
	outputFile := fmt.Sprintf("%s%s/%s", GlobalConfig.Storage.Path, appId, fileHeader.Filename)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		Logger.Infof("creating path: %s", outputPath)
		os.Mkdir(outputPath, os.ModePerm)
	}

	//copy to your file system
	if err := ioutil.WriteFile(outputFile, data, 0666); err != nil {
		Logger.Error(err.Error())
		ResponseBody(w, 90000, "write upload-file error!", "")
		return
	}

	imgUrl := fmt.Sprintf("%s%s/%s", GlobalConfig.Storage.Address, appId, fileHeader.Filename)
	ResponseBody(w, 60000, "success", imgUrl)
}

func DoBase64Upload(w http.ResponseWriter, r *http.Request, p router.Params) {
	//max memory on server - 2MB
	r.ParseForm()
	filename := r.PostForm.Get("filename")
	content := r.PostForm.Get("uploadfile")
	base64Data := strings.Split(content, ",")

	data, err := base64.StdEncoding.DecodeString(strings.Trim(base64Data[1], " "))
	// get file handle
	if err != nil {
		Logger.Error(err.Error())
		ResponseBody(w, 90000, "decode upload-file error!", "")
		return
	}

	// max file size: 300k
	if len(data) > 300<<10 {
		ResponseBody(w, 91000, "maxFileSize more than 300k.", "")
		return
	}

	accepted := false
	for key := range AcceptContentType {
		if strings.Contains(base64Data[0], key) {
			accepted = true
			break
		}
	}

	if !accepted {
		ResponseBody(w, 91000, "only [.jpg .png .gif] supported.", "")
		return
	}

	appId := p.ByName("appId")
	// make sure GlobalConfig.Storage.Path ended with '/'
	outputPath := fmt.Sprintf("%s%s/", GlobalConfig.Storage.Path, appId)
	outputFile := fmt.Sprintf("%s%s/%s", GlobalConfig.Storage.Path, appId, filename)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		Logger.Infof("creating path: %s", outputPath)
		os.Mkdir(outputPath, os.ModePerm)
	}

	//copy to your file system
	if err := ioutil.WriteFile(outputFile, data, 0666); err != nil {
		Logger.Error(err.Error())
		ResponseBody(w, 90000, "write upload-file error!", "")
		return
	}

	imgUrl := fmt.Sprintf("%s%s/%s", GlobalConfig.Storage.Address, appId, filename)
	ResponseBody(w, 60000, "success", imgUrl)
}
