FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s' \
    -o website ./cmd/website

FROM alpine:3.19

RUN apk --no-cache add ca-certificates


COPY --from=builder /build/website /website
COPY --from=builder /build/web /web

EXPOSE 8080

CMD ["/website"]
