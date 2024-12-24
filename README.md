# Auth Service

## Описание

Auth Service — микросервис для управления аутентификацией пользователей с использованием токенов и PostgreSQL для хранения токенов обновления.

## Установка

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/yourusername/auth_service.git
   cd auth_service
   ```

2. Создайте файл `.env` на основе `.example.env` и заполните его данными:

   ```bash
   cp .example.env .env
   ```

3. Установите зависимости:

   ```bash
   go mod download
   ```

## Запуск

### Локально

```bash
go run main.go
```

### С помощью Docker

```bash
docker build -t auth_service .
docker run -p 8080:8080 --env-file .env auth_service
```

## Использование

Сервис доступен по адресу `http://localhost:8080` для управления токенами.