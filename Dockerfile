FROM golang:1.16.3-alpine AS builder

COPY . .

ENV GOPATH=""
RUN go get -d -v ./...
RUN go build cmd/service/main.go

CMD ["./main"]
