FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /app

EXPOSE 8080

ENTRYPOINT ["/app"]