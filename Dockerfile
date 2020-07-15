# builder image
FROM golang:alpine AS builder
WORKDIR /build
COPY . /build
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk add --no-cache git
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOSUMDB=off
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main /build/cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o scheduler /build/cmd/scheduler/main.go


FROM alpine as proxy_pool
COPY --from=builder /build/main  /usr/bin/api
COPY --from=builder /build/scheduler  /usr/bin/scheduler
COPY --from=builder /build/config/conf.yaml /etc/conf.yaml
CMD [ "api" ]