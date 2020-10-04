FROM golang:1.15 as builder

LABEL maintainer="Albert Moreno <albert.moreno.dev@gmail.com>"

WORKDIR /app

COPY ./ /app

RUN go mod download

WORKDIR /app/cmd/todoapi

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todoapi .

######## Start a new stage from scratch #######
FROM alpine:latest

ENV TODOAPI_NAME=todo-api
ENV TODOAPI_HOST=""
ENV TODOAPI_PORT=3000

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/todoapi/todoapi .

EXPOSE ${TODOAPI_PORT}

CMD ["./todoapi"]