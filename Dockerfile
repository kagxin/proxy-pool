# builder image
FROM golang:alpine AS builder
WORKDIR /build
COPY . /build
RUN apk add --no-cache git
RUN go env -w GO111MODULE=on
RUN go env -w GOSUMDB=off
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o proxy-pool /build/exmples/exmple/main.go


FROM alpine as proxy_pool
COPY --from=builder /build/proxy-pool  /usr/bin/proxy-pool