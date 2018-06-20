package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	. "adhoc/adhoc_uploader/auth"
	. "adhoc/adhoc_uploader/logging"
	. "adhoc/adhoc_uploader/handler"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", Auth(SayHello))
	mux.GET("/config/", Auth(Configuration))
	mux.POST("/upload/:appId/", Auth(DoUpload))
	mux.POST("/base64Upload/:appId/", Auth(DoBase64Upload))

	//Catch-all params for HttpOption method.
	mux.OPTIONS("/*path", HttpOption)

	// initialize a server object
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		Logger.Fatal("ListenAndServe: ", err)
	}
}
