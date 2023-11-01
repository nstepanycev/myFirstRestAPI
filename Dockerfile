FROM golang:alpine as development
WORKDIR /app
COPY . .

RUN apk add build-base && go mod download
RUN go install github.com/githubnemo/CompileDaemon@latest

#Seee hot reload from go - https://nuancesprog.ru/p/5657/
RUN go build cmd/url_shortner/main.go
ENTRYPOINT CompileDaemon --build="go build cmd/url_shortner/main.go" --command=./main