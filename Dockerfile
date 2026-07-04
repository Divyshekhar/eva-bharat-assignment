FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .


# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ticket-system ./cmd/main.go
RUN apk add --no-cache gcc musl-dev

RUN go build -o ticket-system ./cmd/main.go



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ticket-system .

EXPOSE 8080

CMD ["./ticket-system"]