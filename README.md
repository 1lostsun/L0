## L0 — ProducerService и OrderService (Kafka demo)

Два Go‑сервиса на Gin:
- ProducerService — принимает заказ по HTTP и публикует в Kafka.
- OrderService — читает из Kafka, сохраняет в Postgres, кеширует в Redis и отдает по HTTP.

Стек: Go, Gin, Kafka+Zookeeper, Postgres 15, Redis 7, Docker Compose.

### Запуск
```bash
docker compose up -d --build
```

### Сервисы и порты
- ProducerService: http://localhost:8081
- OrderService: http://localhost:8082
- Kafka (внутри сети): kafka:9092 (топик: ordersTopic)
- Postgres: order-postgres:5432 (БД: WBOrders, user: lost, pass: 123)
- Redis: order-redis:6379

### Эндпоинты
- POST `/v1/order` (ProducerService) — принимает заказ и отправляет в Kafka
- GET `/v1/order/:order_uid` (OrderService) — возвращает заказ по UID

### Примечания
- Переменные окружения заданы в `docker-compose.yml`.
- Данные Postgres и Redis сохраняются в именованных томах.

### Остановка
```bash
docker compose down -v
```
