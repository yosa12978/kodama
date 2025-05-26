FROM golang:1.24.3-alpine3.21

WORKDIR /app

RUN apk add make
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

CMD [ "air", "-c", ".air.toml" ]
