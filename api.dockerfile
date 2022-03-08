FROM golang:1.17-buster
WORKDIR /app
COPY pkg pkg
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
RUN go build
CMD ["./urlshortener"]