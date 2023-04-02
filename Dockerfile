ARG TELEGRAM_SECRET
ARG GPT_SECRET

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download

COPY . ./

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

EXPOSE 8080

ENTRYPOINT CompileDaemon -polling -log-prefix=false -color=true -build="go build -o joyGPT main.go" -command="./joyGPT"