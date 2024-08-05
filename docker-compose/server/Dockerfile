#FROM golang:1.21.7-alpine
#
#WORKDIR /app
#
## Устанавливаем необходимые инструменты
#RUN apk add --no-cache git
#
## Команда для сборки и запуска приложения
#CMD go mod download && go build -o server ./cmd/server && ./server
#
#

FROM golang:1.21.7-alpine

WORKDIR /app

# Для debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Собираем приложение с отладочной информацией
CMD go mod download && go build -gcflags="all=-N -l" -o server ./cmd/server && ./server

# Команда для запуска приложения в режиме отладки
CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./server"]