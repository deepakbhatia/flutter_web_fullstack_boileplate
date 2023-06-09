FROM golang:1.16.15-alpine3.14

WORKDIR /app

RUN apk add git gcc libc-dev
RUN go install github.com/cosmtrek/air@v1.41.0

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]