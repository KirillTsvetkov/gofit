# Используем официальный образ Golang как базовый
FROM golang:1.21

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod ./
COPY go.sum ./

# Скачиваем все зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию
COPY . .

# Переходим в подкаталог с исходным кодом
WORKDIR /app/cmd/gofit

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Переходим обратно в корневую директорию
WORKDIR /app

# Открываем порт  8080
EXPOSE  8080

# Запускаем приложение
CMD ["/app/cmd/gofit/main"]