## gotti-server

Простой сервер для хранения [lottie](https://lottiefiles.com/) анимаций

Стек:
- [gofiber](https://github.com/gofiber/fiber)
- [mongodb](https://www.mongodb.com/)
- [swagger](https://swagger.io/)
- [prometheus](https://prometheus.io/)
- [grafana](https://grafana.com/)
- [alertmanager](https://github.com/prometheus/alertmanager)
- [cadvisor](https://github.com/google/cadvisor)

## Запустить локально

1. Запустить mongodb
2. Отредактировать .env.local
3. run:
```
go run main.go
```

- gotti-server http://127.0.0.1:9000/api/gotti
- метрики http://localhost:9000/metrics
- swagger http://127.0.0.1:9000/swagger/index.html

## Запустить через docker-compose

дополнительно будут доступны метрики: prometheus grafana alertmanager cadvisor

1. Переименовать файл .env.local (приложение сначала читает этот файл, потом .env)
2. Отредактировать .env
3. run:
```
docker-compose up --build
```

- prometheus http://localhost:9090/targets
- grafana http://localhost:3000/ 
- alertmanager http://localhost:9093
- cadvisor http://localhost:8080
