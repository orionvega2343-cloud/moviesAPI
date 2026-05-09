#  Movies Review API

> Rest API для просмотра и получения фильмов с возможностью оставлять отзывы, регистрироваться и получать список просмотров

**Стек:** Go, Chi, PostgreSQL, JWT, Bcrypt, godotenv, cleanenv, sqlx, golang-migrate, Docker

---

## Содержание

- [Установка и запуск](#установка-и-запуск)
- [Создание .env файла](#создание-env-файла)
- [Запуск БД через Docker](#запуск-бд-через-docker)
- [Применение миграций](#применение-миграций)
- [Запуск сервера](#запуск-сервера)
- [Эндпоинты API](#эндпоинты-api)
- [Переменные окружения](#переменные-окружения)
- [Структура проекта](#структура-проекта)

---

## Установка и запуск

Клонируйте репозиторий на ваш компьютер:

```bash
git clone https://github.com/orionvega2343-cloud/moviesAPI.git
cd moviesAPI
```

Установите зависимости:

```bash
go mod tidy
```

---

## Создание .env файла

В корне проекта создайте файл `.env` и заполните его следующими переменными:

```env
DB_PASS=ваш_пароль_от_базы_данных
JWT_SECRET=любая_длинная_строка_символов
```

>  Никогда не добавляйте `.env` в git. Убедитесь что он есть в `.gitignore`

---

## Запуск БД через Docker

Запустите PostgreSQL контейнер с помощью docker-compose:

```bash
docker-compose up -d
```

Проверьте что контейнер запущен:

```bash
docker ps
```

---

## Применение миграций

Установите утилиту golang-migrate (если ещё не установлена):

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Примените миграции:

```bash
migrate -path migrations -database "postgres://postgres:ПАРОЛЬ_ОТ_БД@127.0.0.1:5431/movies?sslmode=disable" up
```

Для отката миграций:

```bash
migrate -path migrations -database "postgres://postgres:ПАРОЛЬ_ОТ_БД@127.0.0.1:5431/movies?sslmode=disable" down
```

---

## Запуск сервера

```bash
go run cmd/main.go
```

Сервер запустится на порту указанном в `config.yml` (по умолчанию `:8080`)

---

## Эндпоинты API

###  Публичные маршруты

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/auth/register` | Регистрация нового пользователя |
| `POST` | `/auth/login` | Авторизация, возвращает JWT токен |

**Пример запроса на регистрацию:**
```json
POST /auth/register
{
  "email": "user@example.com",
  "name": "Иван",
  "password": "секретный_пароль"
}
```

**Пример запроса на логин:**
```json
POST /auth/login
{
  "email": "user@example.com",
  "password": "секретный_пароль"
}
```

---

###  Защищённые маршруты

> Для доступа необходимо передать JWT токен в заголовке:
> `Authorization: Bearer <ваш_токен>`

#### Фильмы

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/movie` | Создание нового фильма |
| `GET` | `/movie/{id}` | Получение фильма по ID |

**Пример запроса на создание фильма:**
```json
POST /movie
{
  "title": "Inception",
  "year": 2010,
  "genre": "Sci-Fi",
  "director": "Christopher Nolan"
}
```

#### Отзывы

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/reviews` | Создание отзыва на фильм |
| `GET` | `/reviews/{id}` | Получение отзыва по ID |

**Пример запроса на создание отзыва:**
```json
POST /reviews
{
  "movie_id": 1,
  "user_id": 1,
  "rating": 9,
  "text": "Отличный фильм!"
}
```

>  Рейтинг должен быть от 1 до 10

#### Список просмотров

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/watchlist` | Добавление фильма в список просмотров |
| `GET` | `/watchlist/{id}` | Получение записи из списка по ID |

---

## Переменные окружения

| Переменная | Описание | Пример |
|------------|----------|--------|
| `DB_PASS` | Пароль от базы данных PostgreSQL | `mypassword` |
| `JWT_SECRET` | Секретный ключ для подписи JWT токенов | `supersecretkey123` |

---

## Структура проекта

```
moviesAPI/
├── cmd/
│   └── main.go              # Точка входа
├── Internal/
│   ├── config/              # Конфигурация приложения
│   ├── db/                  # Подключение к БД
│   ├── handlers/            # HTTP хэндлеры
│   ├── middlewares/         # Auth и Logger мидлвары
│   ├── models/              # Модели данных
│   ├── repository/          # Слой работы с БД
│   └── service/             # Бизнес логика
├── migrations/              # SQL миграции
├── config.yml               # Конфигурационный файл
├── docker-compose.yml       # Docker конфигурация
└── .env                     # Переменные окружения (не в git)
```
