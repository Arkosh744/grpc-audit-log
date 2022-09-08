# Сервис логирования событий для CRUD APP

### Стэк

- go 1.18
- mongoDB

### Запуск

Для запуска необходимо указать переменные окружения, например в файле .env директории internal/config/app.env

```
DB_URI=mongodb://localhost:27017
DB_USERNAME=admin
DB_PASSWORD=g0langn1nja
DB_DATABASE=audit
SERVER_PORT=9000
```

Для mongo используем Docker

```
docker run --rm -d --name audit-log-mongo -e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=g0langn1nja -p 27017:27017 mongo:latest
```