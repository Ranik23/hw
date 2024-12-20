create-db:
	psql -c "CREATE DATABASE library;" || echo 'Database already exists, skipping creation.'
	psql -d library -c "CREATE SCHEMA IF NOT EXISTS library;" || echo 'Schema already exists, skipping creation.'
	psql -d library -c "SET search_path TO '\$$user', library, public;" || echo 'Failed to set search path.'

migrate-up: create-db
	go run cmd/migrator/main.go -command=up

migrate-down:
	go run cmd/migrator/main.go -command=down


run:
	go run cmd/main/main.go