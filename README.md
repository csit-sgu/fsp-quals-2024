# zero-cost developers

## Общая идея

Сервис на Python тянет данные из PDF, парсит их и сохраняем в Clickhouse.

При получении запроса от клиентская часть приложения создаёт запрос к Clickhouse и возвращает результат.

Вся логика фильтрации реализована через формирование динамического запроса к представлению в Clickhouse.

### Фильтрация по грязным данным

Так как PDF может содержать некоторые грязные данные, то при фильтрации по грязным данным необходимо обрабатывать их как текст.

Для этого планируется индексировать текст о дисциплинах и соревнованиях и выдавать наиболее релевантные результаты.

## Стек технологий

- Python
- Clickhouse
- Golang
- Vue.js
- Docker Compose

## Архитектура

![Архитектура](docs/arch.png)

## Текущее состояние проекта

![Текущее состояние](docs/state.png)

## Запуск проекта

### Backend

Все сервисы бэкенда запускаются

```shell
docker compose -f deploy/compose.yaml up -d
```

### Frontend

Есть возможность запустить проект в режиме разработки. Для этого необходимо перейти в каталог
frontend и запустить `npm run dev`. После этого можно перейти на страницу
`http://localhost:5173`. Перед этим необходимо указать переменную
`VITE_BACKEND_URL=http://localhost:3000` в файле `.env.local`.
