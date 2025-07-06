FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod download

# Ключевое изменение - статическая компиляция
RUN CGO_ENABLED=0 go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

# Обеспечивает совместимость, если бинарник все же требует каких-то библиотек
RUN apk add --no-cache libc6-compat

# Берем файл из образа-стадии builder, /app/main - путь до бинарника, кидаем в текущую рабочую директорию
COPY --from=builder /app/main .

# Документируем работу на порту 8080
# При меппинге чтобы контейнер работал docker run -p порт_хоста:8080 my-image, только так 
EXPOSE 8080

CMD ["./main"]