FROM golang:1.17-alpine AS build
RUN apk add --no-cache git
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
COPY httpserver.go ./
RUN go mod download
RUN go build -o httpserver .

# FROM scratch
# COPY --from=build /build/httpserver /
ENTRYPOINT ["./httpserver"]
# CMD ["--help"]