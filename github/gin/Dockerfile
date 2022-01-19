FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go build -o app

EXPOSE 80
CMD ["./app"]