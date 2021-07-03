FROM golang:latest

WORKDIR /go/apiBeatsPro

COPY ./ /go/apiBeatsPro

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/app/main.go" --command=./main