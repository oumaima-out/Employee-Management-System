
FROM golang:1.23.3-alpine3.20

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download && \
  go mod verify

RUN go build -o survey .

CMD ["./survey"]

EXPOSE 8080

