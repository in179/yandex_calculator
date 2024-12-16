# Сервис Калькулятор

Этот проект представляет собой веб-сервис для вычисления математических выражений, отправляемых через HTTP.

## Возможности

- Принимает математические выражения в POST-запросах.
- Возвращает вычисленный результат или соответствующее сообщение об ошибке.
- Три возможных сценария ответа:
  - Успешное вычисление (HTTP 200)
  - Неверное выражение (HTTP 422)
  - Внутренняя ошибка сервера (HTTP 500)

## Спецификация API

### Эндпоинт

`POST /api/v1/calculate`

### Формат запроса

```json
{
  "expression": "ваше_выражение"
}
```

### Формат ответа

Успешное вычисление (HTTP 200)
```json
{
  "result": "вычисленный_результат"
}
```

### Неверное выражение (HTTP 422)
```json
{
  "error": "Expression is not valid"
}
```
### Внутренняя ошибка сервера (HTTP 500)
```json
{
  "error": "Internal server error"
}
```
## Примеры использования
### Успешный запрос
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
### Ответ:
```json
{
  "result": "6.000000"
}
```
### Неверное выражение
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 / 0"
}'
```
### Ответ:
```json
{
  "error": "Expression is not valid"
}
```
### Внутренняя ошибка сервера

### Если произошла неожиданная ошибка:
```json
{
  "error": "Internal server error"
}
```
## Установка и запуск

### Клонируйте репозиторий:

```bash
git clone https://github.com/in179/yandex_calculator.git
```

### Перейдите в папку проекта:
```bash
cd yandex_calculator
```
### Запустите сервис:
```bash
go run ./main.go
```
## Сервис будет доступен по адресу http://localhost:8080.
