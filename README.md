# GoalTracker

Простой учебный pet-проект на `go` и `React`. Структура:
* `./pkg` — бек
* `./website` — фронт

## Конфигурация
Для начала, необходим сделать `.env` файл:
```sh
cp .env.example .env
```

Чтобы работал счетчик объявлений для [auto.ru](https://auto.ru), нужно подставить куки из браузера в переменную окружения `APP_AUTORU_COOKIE`.

## Запуск
### В `Docker`:
```sh
docker-compose up
```

### Запуск в режиме разработки:
**Back:**
```sh
docker-compose up -d db
source .env
cd ./pkg && air
```
Либо: `task run`

Генерация/актуализация swagger-документации
```
swag init -g ./cmd/main/main.go --dir ./ --output ./docs
```

**Front:**
```sh
cd ./website && npm run start
```

## Task-runner
В качестве task runner в проекте используется [go-task](https://taskfile.dev). [Установка](https://taskfile.dev/installation/) и использование:
```sh
brew install go-task
task TASK_NAME
```

### Доступные команды
* `run` запуск бекенда
* `db` консоль базы данных
* `drop_db` удаление базы данных
  
## @todo
- [x] Прикрутить релоадер для режима разработки
- [x] Напиcать докер-файлы
- [x] Сделать отделные модели для request/response
- [x] Привести структуру проекта в привычную для go
- [x] Прикрутить редис и кешировать счетчик там 
- [x] Считать правильную сумму по накоплениям
- [ ] Написать тесты

### Ideas
- Прикрутить авторизацию
- Сделать хэндлеры на интерфейсах
- 