FROM golang:latest

WORKDIR /go/apiBeatsPro

COPY . /go/apiBeatsPro

RUN export GO111MODULE="on" && \
    chmod +x scripts/ -R && \
    apt update && \
    apt install -y curl && \
    go mod download && \
    go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/app/main.go" --command=./main