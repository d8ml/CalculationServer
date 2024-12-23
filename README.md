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

Запуск unit-тестов с покрытием
```
go test -cover ./calc/internal/pkg/calculation/
```

Пример
```
curl --location 'http://localhost:8082/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "1 + 2"
}'
```
Ответ 200 OK
```
{
    "result": 3
}
```

Некорректное выражение
```
curl --location 'http://localhost:8082/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "1 + (2"
}'
```
Ответ 422 Unprocessable Entity
```
{
    "error": "Expression is not valid"
}
```
Интеграционные тесты находятся в папке calc/test

Unit-тесты находятся в папке calc/internal/pkg/calculation

Запуск всех тестов
```
go test ./...
```