FROM golang:latest

WORKDIR /app
COPY .env ./
ENTRYPOINT ["go", "run", "cmd/main.go"]