FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app
RUN apk add --no-cache make
COPY . .
RUN make

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/bin/* .
RUN apk add --no-cache curl
EXPOSE 80
ENTRYPOINT [ "/app/kodama" ]
