FROM golang:latest

WORKDIR /BeatsPro

COPY ./ /BeatsPro

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/app/main.go" --command=./main