FROM golang:1.22-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o quackbin ./cmd/server

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/quackbin .

COPY web/ ./web/

EXPOSE 8080

CMD ["./quackbin"]
