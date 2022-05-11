FROM golang:1.17-buster

RUN mkdir /build
WORKDIR /buid
COPY . .
RUN make bin

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait for postgres executable
RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o trade-bot ./cmd/api/main.go
RUN go build -o trade-bot-client ./pkg/telegramBot/cmd/api/main.go

CMD ["./trade-bot"]
CMD ["./trade-bot-client"]

