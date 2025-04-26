FROM golang:1.24.2-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -tags netgo -ldflags '-s -w' -o app

CMD ["./app"]