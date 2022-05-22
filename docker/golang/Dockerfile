FROM golang:1.18.2-alpine

ENV GO111MODULE=on

WORKDIR /go/src/api

COPY ./api /go/src/api

RUN apk update && \
    apk upgrade && \
    apk --no-cache add git && \
    go install github.com/cosmtrek/air@v1.27.3

RUN go mod tidy
RUN go mod download

CMD ["air", "-c", ".air.toml"]