FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY . .

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go build -o sso ./cmd/sso/main.go

FROM alpine:latest

COPY --from=builder /app/sso /sso
COPY ./swagger/swagger.json ./swagger/swagger.json
COPY ./.env .

CMD ["/sso"]
