FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

EXPOSE 8080

ENTRYPOINT CompileDaemon -polling -log-prefix=false -color=true -build="go build -o cmd/app/bin/joyGPT cmd/app/main.go" -command="./cmd/app/bin/joyGPT"