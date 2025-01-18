FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY vendor /app/vendor
COPY go.mod .
COPY go.sum .
COPY . .

RUN wget https://github.com/upx/upx/releases/download/v4.2.1/upx-4.2.1-amd64_linux.tar.xz && \
    tar -xf upx-4.2.1-amd64_linux.tar.xz && \
    mv upx-4.2.1-amd64_linux/upx /usr/local/bin/ && \
    rm -rf upx-4.2.1-amd64_linux.tar.xz upx-4.2.1-amd64_linux

RUN CGO_ENABLED=0 go build -o /file-map-server ./cmd/main.go
RUN upx --best --lzma /file-map-server

FROM alpine:3.21 AS runtime
WORKDIR /app

COPY --from=builder /file-map-server /app/file-map-server

COPY ./app/config/config.yaml /app/config.yaml
COPY ./app/dist /app/dist

ENTRYPOINT ["./file-map-server"]
