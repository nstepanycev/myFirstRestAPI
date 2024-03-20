FROM golang:alpine as local
WORKDIR /app
COPY . .

RUN apk add build-base && go mod download
RUN go install github.com/githubnemo/CompileDaemon@latest

#Seee hot reload from go - https://nuancesprog.ru/p/5657/
RUN go build cmd/main.go
ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command=./main