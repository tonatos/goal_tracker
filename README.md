# GoalTracker

Простой учебный pet-проект на `go` и `React`. Структура:
* `./api` — бек
* `./web` — фронт


## Запуск
### В `Docker`:
```
cp .env.example .env
docker-compose up
```

### Запуск в режиме разработки:
Back:
```
docker-compose up -d db
cp .env.example .env
source .env
cd ./api && go run ./main.go
```
Либо: `task run`

Front:
```
cd ./web && npm run start
```

## Task-runner
В качестве task runner в проекте используется [go-task](https://taskfile.dev). [Установка](https://taskfile.dev/installation/) и использование:
```
brew install go-task
task TASK_NAME
```

### Доступные команды
* `run` запуск бекенда
* `db` консоль базы данных
* `drop_db` удаление базы данных