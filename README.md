# GoalTracker

Простой учебный pet-проект на `go` (Fiber) и `TypeScript` (React). Структура:
* `./build` — докер-файл
* `./cmd` — инициализация приложения
* `./docs` — сваггер-схема
* `./internal` — доменный код
* `./pkg` — код контура окружения и утилиты
* `./test` — тесты бекенда
* `./website` — фронт

Основано на [project-layout](https://github.com/golang-standards/project-layout).

## Конфигурация
Для начала, необходим сделать `.env` файл:
```sh
cp .env.example .env
source .env
```

Чтобы работал счетчик объявлений для [auto.ru](https://auto.ru), нужно подставить куки из браузера в переменную окружения `API_AUTORU_COOKIE`.

## Запуск
### В `Docker`:
```sh
docker-compose up
```

Swagger будет доступен по ссылке [http://127.0.0.1:8000/api/v1/docs/index.html](http://127.0.0.1:8000/api/v1/docs/index.html).

### Запуск в режиме разработки:
#### Backend:
**Запуск:**
```sh
docker-compose up -d db redis
air
```
Либо: `task run_api`

**Генерация/актуализация swagger-документации:**
```sh
swag init -g ./cmd/main/main.go --dir ./ --output ./docs
```

**Тесты:**
```sh
docker-compose up -d db redis
export $(cat .env) | ENV=test go test ./test/... -count=1 -v
```
Либо: `task test_back`

#### Frontend:
**Установка:**
```sh
cd ./website && yarn install
```

**Запуск:**
```sh
npm run start
```
Либо: `task run_front`

## Task-runner
В качестве task runner в проекте используется [go-task](https://taskfile.dev). [Установка](https://taskfile.dev/installation/) и использование:
```sh
brew install go-task
task TASK_NAME
```

### Доступные команды
* `run_api` запуск бекенда
* `run_front` запуск фронта
* `test_api` локальные тесты
* `db_shell` консоль базы данных
* `drop_db` удаление базы данных
  
## @todo
- [x] Прикрутить релоадер для режима разработки
- [x] Напиcать докер-файлы
- [x] Сделать отделные модели для request/response
- [x] Привести структуру проекта в привычную для go
- [x] Прикрутить редис и кешировать счетчик там 
- [x] Считать правильную сумму по накоплениям
- [x] Написать тесты
  - [x] Протестировать API-ручки
  - [x] Замокать АПИху авто-ру
  - [x] Протестировать сервисы
- [ ] Упаковать фронт в докер и добавить traefik
- [ ] Сделать выбор целей / добавление целей
- [ ] Сделать редактирование платежей (круд вокруг них)
- [ ] Заверстать мобилку
- [ ] Написать тесты на фронте

### Ideas
- Прикрутить авторизацию
- Сделать хэндлеры на интерфейсах
- Сделать нормально конфигурируемые сеттинги
- Прикрутить парсер авито/дром, парсить их параллельно через горутины и аггрегировать результат
