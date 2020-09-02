FROM golang:1.15 as builder

LABEL maintainer="Albert Moreno <albert.moreno.dev@gmail.com>"

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o server cmd/server/main.go" --command=./server