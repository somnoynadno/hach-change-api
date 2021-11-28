# Hack&Change 2021

Исходный код серверной части проекта (BFF).

## Модель предметной области

![mpo](http://somnoynadno.ru/static/data/hack.png)

## Деплой проекта

Проект состоит из трёх сервисов (см. [спеку](https://github.com/somnoynadno/hack-change-api/blob/master/docker-compose.yml)):
- **invest-api** - непосредственно APIшка на Go
- **invest-db** - PostgreSQL для неё
- **invest-adminer** - веб-интерфейс для БД

Чтобы поднять сервисы в контейнере: ` $ docker-compose up --build -d`

## Документация API

Доступна в формате Postman: https://www.getpostman.com/collections/924501d84f116cafa378

Создание придерживалось идеологии REST API.

Здесь можно посмотреть список [моделей](https://github.com/somnoynadno/hack-change-api/tree/master/models/entities) и 
[эндпоинтов](https://github.com/somnoynadno/hack-change-api/tree/master/server/api) сервиса.

## Дополнительно

По всем вопросам: [@somnoynadno](https://t.me/somnoynadno)
