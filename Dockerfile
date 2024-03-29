FROM golang:latest

WORKDIR /app/cmd

ENTRYPOINT ["go", "run", "."]