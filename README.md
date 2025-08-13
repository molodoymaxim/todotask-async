# Task API

Простой REST API на Go для управления задачами с in-memory хранилищем, асинхронным логированием и graceful shutdown.

## Функционал
- GET /tasks?status="status" — Получить список задач (фильтрация по статусу)
- GET /tasks/{id} — Получить задачу по ID
- POST /tasks — Создать задачу (JSON: {"title": "string", "status": "string"})

## Сборка и запуск
### 1. Клонирование репозитория
   ```bash
   git clone <repository_url>
   cd todotask-async
   ```
### 2. Настройка .env
Создайте файл `.env` в корне проекта на основе `.env.example`:
```plaintext
PORT=1234
```
### 3. Установка зависимостей
```bash
go mod tidy
```
### 3. Запуск Docker
```bash
docker-compose build && docker-compose up
```

## Тестирование
Примеры тестовых запросов:
- `GET /tasks`
- `GET /tasks?status=<some_status>`
- `GET /tasks/{id}`
- `POST /tasks` с `{"title": "Write document", "status": "todo"}`

## Логирование
Асинхронное через канал, вывод в stdout: "Action: <действие>, <детали>".
