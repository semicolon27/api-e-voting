FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./app

ENTRYPOINT ["./app"]