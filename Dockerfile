FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /read-adviser-tgbot

EXPOSE 8080

CMD ["/read-adviser-tgbot", "-tg-bot-token=${TELEGRAM_API_TOKEN}"]