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

(POST}:
`http://localhost:8080/auth/tokens`


Params:

```
   Key: user_id
   Value: string_value
```

Вернёт JSON
```json
   {
      "access_token": "access_token_value",
       "refresh_token": "refresh_token_value"
   }
```
(POST):
`http://localhost:8080/auth/refresh`


Body (raw JSON): 
```json
   {
      "access_token": "your_access_token_value",
      "refresh_token": "your_refresh_token_value"
   }
```

Вернёт JSON
```json
   {
      "access_token": "new_access_token_value",
      "refresh_token": "new_refresh_token_value"
   }
```
