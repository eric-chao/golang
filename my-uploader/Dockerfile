FROM golang:1.9-alpine3.6

MAINTAINER ZXJ

WORKDIR /go/src/adhoc/adhoc-uploader
COPY . /go/src/adhoc/adhoc-uploader

RUN echo 'https://mirrors.aliyun.com/alpine/v3.7/main/' > /etc/apk/repositories \
    && echo 'https://mirrors.aliyun.com/alpine/v3.7/community/' >> /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' > /etc/timezone \
    && apk del tzdata \
    && mv /go/src/adhoc/adhoc-uploader/vendor/* /go/src/

ENV LANG en_US.utf8

RUN go build .

CMD ["./adhoc-uploader"]