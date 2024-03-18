FROM golang:1.22.1-alpine3.19 as builder
WORKDIR /build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache make git && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /build/bin/app /app/

ENTRYPOINT ["/app/app"]