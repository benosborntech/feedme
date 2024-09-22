FROM golang:1.23.1 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app apigw/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/app .

CMD ["./app"]