# README.md

## Описание проекта

Этот проект предназначен для управления базой данных библиотеки с использованием PostgreSQL и миграций на Go. В этом файле описаны команды для создания базы данных, схемы, а также для выполнения миграций и запуска приложения.

## Установка

Перед началом работы убедитесь, что у вас установлены следующие компоненты:

- PostgreSQL
- Go (версия 1.16 или выше)

## Команды

### Создание базы данных

```bash
psql -c "CREATE DATABASE library;" || echo 'Database already exists, skipping creation.'
```

### Создание схемы

```bash
psql -d library -c "CREATE SCHEMA IF NOT EXISTS library;" || echo 'Schema already exists, skipping creation.'
```
### Установка пути поиска

```bash
psql -d library -c "SET search_path TO '\$$user', library, public;" || echo 'Failed to set search path.'
```

### Миграция UP
```bash
go run cmd/migrator/main.go -command=up
```

### Миграция DOWN
```bash
go run cmd/migrator/main.go -command=down
```
