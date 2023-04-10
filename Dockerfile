FROM golang:alpine as builder

LABEL maintainer="pratheesh pc <pratheeshpcpalappetta@gmail.com>"

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/
COPY *.go /app/

RUN go mod download

RUN go build -o cmd /app/main.go


# Expose the port that the echo service listens on
EXPOSE 8080

# Start the service
CMD ["./cmd"]
