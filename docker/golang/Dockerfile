FROM golang:1.21-alpine

RUN apk add --update && apk add git

# for Vscode extension
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest

WORKDIR /app
