FROM golang:1.19-alpine

# Установка необходимых зависимостей для CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Установка переменной окружения для включения CGO
ENV CGO_ENABLED=1

# Копируем файлы go.mod и go.sum и загружаем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Установка токена в переменной окружения
ENV TELEGRAM_API_TOKEN=7752693039:AAGEMKLZYj7mm-z8TcFJ7rOhC0AIfhdsU6s

# Сборка проекта
RUN go build -o /read-adviser-tgbot

EXPOSE 8080

CMD ["/read-adviser-tgbot", "-tg-bot-token=${TELEGRAM_API_TOKEN}"]