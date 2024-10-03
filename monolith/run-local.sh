#!/bin/bash
# db cache network
docker-compose -f docker-compose-local.yml up -d
# Собираем приложение
go build -o ./temp/local/main ./cmd/api/main.go
cp ./.env ./temp/local
go run ./cmd/api/main.go
rm -R ./temp