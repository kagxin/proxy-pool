# builder image
FROM golang:alpine AS builder
WORKDIR /build
COPY . /build
RUN apk add --no-cache git
RUN go env -w GO111MODULE=on 
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main /build/cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o scheduler /build/cmd/scheduler/main.go


FROM alpine as proxy_pool
COPY --from=builder /build/main  /usr/bin/api
COPY --from=builder /build/scheduler  /usr/bin/scheduler
COPY --from=builder /build/config/conf.yaml /etc/conf.yaml