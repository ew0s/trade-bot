FROM golang:1.17-buster

RUN mkdir /build
WORKDIR /build
COPY . .
RUN make bin

RUN go build -o trade-bot-api ./cmd/api/main.go

CMD ["./trade-bot-api"]

