# CalculationServer

Запуск сервера c портом по умолчанию(8082)
```
go run ./calc/cmd/server/server.go 
```

Запуск сервера c заданием порта 
```
export PORT=8083 && go run ./calc/cmd/server/server.go
```

Сервер пишет логи в консоль

Для обращения к серверу
```
curl --location 'http://localhost:8083/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "22 / 2"
}'
```