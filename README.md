# Task L0
```
Запуск:
    make build
Сервер доступен локально по адресу:
    http://localhost:8080/orders/
Для публикации данных в канал nats-streaming:
    go run publisher.go

Api c подключением и подпиской на канал в nats-streaming, данные записываются в Postgres и хранятся in memory (кеш).
Кеш восстанавливается из Postgres. Http сервер выдает данные по id из кеша.
