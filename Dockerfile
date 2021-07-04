##
## Build
##

FROM golang:1.16-buster AS build

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /golang-demo

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /golang-demo /golang-demo

ENTRYPOINT ["/golang-demo"]