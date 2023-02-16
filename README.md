# GoalTracker

Простой учебный pet-проект на `go` и `React`. Структура:
* `./api` — бек
* `./api` — фронт


## Запуск
### В `Docker`:
```
docker-compose up
```

### Запуск в режиме разработки:
```
docker-compose up -d db
cd ./api && got run ./main.go # запуск серверной части
cd ./web && npm run start # запуск фронт
```