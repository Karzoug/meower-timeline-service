# Meower timeline service

Сервис ответственный за формирование ленты пользователя.

Предоставляет доступ к ленте пользователя посредством grpc c интерфейсом и сообщениями описанным в [api](https://github.com/Karzoug/meower-api/tree/main/proto/timeline). Сервис получает события только из своего топика kafka. Преобразование событий других сервисов в события сервиса ленты осуществляются [pipeline сервисом](https://github.com/Karzoug/meower-timeline-pipeline) (включая основной fan-out сценарий).

### Стек
- Основной язык: go
- База данных: redis
- Брокер: kafka
- Наблюдаемость: opentelemetry, jaeger, prometheus
- Контейнеры: docker, docker compose

## Дальнейшее развитие

- [ ] решение "проблемы знаменитостей"
- [ ] кастомные метрики
- [ ] тесты