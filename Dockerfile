FROM golang:1.20.1-alpine


WORKDIR /binance-ticker-parser

COPY . .

RUN go mod download

RUN go build -o parser ./cmd/parser

EXPOSE 4000

CMD ["./parser"]
