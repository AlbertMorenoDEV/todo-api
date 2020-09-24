FROM golang:1.15 as builder

LABEL maintainer="Albert Moreno <albert.moreno.dev@gmail.com>"

WORKDIR /app

COPY ../../../../go.mod go.sum ./

RUN go mod download

COPY ../../../.. .

WORKDIR /app/cmd/todoapi

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todoapi .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/todoapi/todoapi .

EXPOSE 8080

CMD ["./todoapi"]