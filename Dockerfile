FROM golang:1.24.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN go mod tidy
RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
