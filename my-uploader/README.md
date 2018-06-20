# FileUpload
Basic File Upload Using Go 


Steps to follow
```
git clone
cd FileUpload
go run main.go

Access 127.0.0.1:8080/upload
```

[References](https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.5.html)

[docker]
0. godep installation
go get github.com/tools/godep

1. run godep save [save dependency to vendor]
$GOPATH/bin/godep save

[docker test]
docker run --rm -it -p 8080:8080 --name upload-local registry.appadhoc.com:30443/adhoc-uploader:dev