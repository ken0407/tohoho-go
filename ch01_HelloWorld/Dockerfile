FROM golang:latest
COPY sample.go .
RUN gofmt sample.go
RUN go build -o ./helloworld ./sample.go
