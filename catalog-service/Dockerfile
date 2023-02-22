FROM golang:1.19.3-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod tidy
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /main ./cmd

FROM alpine:3.16.3

COPY --from=builder /main .

CMD ["./main"]

