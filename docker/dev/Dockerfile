FROM golang:1.24-alpine AS dev

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

CMD ["CompileDaemon", "--build=go build -o ./tmp/website ./cmd/website", "--command=./tmp/website", "--color=true"]
