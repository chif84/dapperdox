Title: Введение в Web API
SortOrder: 100

# Введение в Web API
API ELMA365 — это инструментарий для интеграции ELMA365 со сторонними системами.

## Общая информация
- Всё API работает по протоколу HTTPS, путем выполнения POST-запросов.
- Авторизация осуществляется по токену.
- Все данные доступны только в формате JSON.
- Базовый URL - https://{company}.elma365.ru/pub/v1.
- Возможны запросы к данным из любой системы.

>Размер веб запроса не может превышать 50 мегабайт

## Авторизация по токену
- Токен передается в заголовке http-запроса.

`«X-Token»: «4bece94a-d372-404f-a4be-fb159019d9e0»`

- Для пользователя индивидуально создается токен и все запросы отправляются от имени пользователя.

Управление правами на API осуществляется с помощью токена.