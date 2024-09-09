# Guests API

## Введение

Проект реализует API для CRUD операций над сущностью "Гость". То есть принимает данные для создания, изменения,
получения, удаления записей гостей хранящихся в выбранной базе данных.

## Установка

1. `cp .env.example .env`
2. По необходимости, отредактировать данные в `.env`
3. `docker-compose up -d`

## Описание методов API

### Получить список всех гостей

- Метод GET
- Путь: `/guests`
- Пример ответа:

```json
{
  "status": "OK",
  "code": 200,
  "message": "The list of all users.",
  "data": {
    "guests": [
      {
        "id": 1,
        "name": "Jacques",
        "last_name": "Webster",
        "email": "jacqueswebster@gmail.com",
        "phone": "+1(555)555-1234",
        "country": "US"
      },
      {
        "id": 2,
        "name": "Jordan",
        "last_name": "Carter",
        "email": "jordancarter@gmail.com",
        "phone": "+1(555)555-5678",
        "country": "US"
      },
      {
        "id": 3,
        "name": "Oleg",
        "last_name": "Dinkov",
        "email": "dinkovv@gmail.com",
        "phone": "+7(952)812-52-52",
        "country": "RU"
      }
    ]
  }
}
```

### Получить пользователя
- Метод GET
- Путь: `/guest/{id}`
- Пример ответа:

```json
{
  "status": "OK",
  "code": 200,
  "message": "The user was successfully retrieved.",
  "data": {
    "guest": {
      "id": 1,
      "name": "Jacques",
      "last_name": "Webster",
      "email": "jacqueswebster@gmail.com",
      "phone": "+1(555)555-1234",
      "country": "US"
    }
  }
}
```

- Ошибки:
    - 404 Not Found

### Создать пользователя
- Метод POST
- Путь `/guest`
- x-www-form-url-encoded:
    - ```
      "name": "Jacques",
      "last_name": "Webster",
      "email": "jacqueswebster@gmail.com",
      "phone": "+1(555)555-1234",
      "country": "US"
    
    - ```
      "name": "Oleg",
      "lastName": "Dinkov",
      "email": "dinkovv@gmail.com",
      "phone": "+7(952)812-52-52"
- Пример ответа
```json
{
    "status": "Created",
    "code": 201,
    "message": "The user was successfully created.",
    "data": {
        "guest": {
            "id": 3,
            "name": "Oleg",
            "lastName": "Dinkov",
            "email": "dinkovv@gmail.com",
            "phone": "+7(952)812-52-52",
            "country": "RU"
        }
    }
}
```
- Ошибки:
  - 400 Bad Request

### Обновить пользователя
- Метод PATCH
- Путь `/guest`
- Пример x-www-form-url-encoded:
    - ```
      "id": 1
      "email": "jacqueswebstersheesh@gmail.com",
- Пример ответа:
```json
{
    "status": "Created",
    "code": 200,
    "message": "The user was successfully updated.",
    "data": {
        "guest": {
          "name": "Jacques",
          "last_name": "Webster",
          "email": "jacqueswebstersheesh@gmail.com",
          "phone": "+1(555)555-1234",
          "country": "US"
        }
    }
}
```
- Ошибки:
  - 404 Not Found
  - 400 Bad Request

### Удалить пользователя

- Метод DELETE
- Путь: `/guest/{id}`
- Пример ответа:
```json
{
  "status": "OK",
  "code": 200,
  "message": "The user was successfully deleted.",
  "data": {
    "result": 1
  }
}
```
- Ошибки:
  - 400 Bad Request

### Формат ошибок
```json
{
    "error": {
        "status": "Not Found",
        "code": 404,
        "message": "id: the id value 1 was not found"
    }
}
```