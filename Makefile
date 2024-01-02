.PHONY: migrate_force migrate_version migrate_down migrate_up new_migration local run build test

DB_URL=postgres://root:root@localhost:5432/pomogo?sslmode=disable
MIGRATION_PATH=./configs/migrations

# ==============================================================================
# Go migrate postgresql

migrate_force:
	migrate -database "$(DB_URL)"  -path "$(MIGRATION_PATH)" force 1

migrate_version:
	migrate -database "$(DB_URL)" -path "$(MIGRATION_PATH)" version

migrate_up:
	migrate -database "$(DB_URL)" -path "$(MIGRATION_PATH)" up 1

migrate_down:
	migrate -database "$(DB_URL)" -path "$(MIGRATION_PATH)" down 1

new_migration:
	migrate create -ext sql -dir "$(MIGRATION_PATH)" -seq $(name)


# ==============================================================================
# Docker compose commands

local:
	echo "Starting local environment"
	docker-compose -f ./deployments/docker-compose.yml up -d


# ==============================================================================
# Tools commands

swaggo:
	echo "Starting swagger generating"
	swag init -o ./api -g **/**/*.go

swaggo_fmt:
	echo "Starting format the SWAG comments"
	swag fmt


# ==============================================================================
# Main

run:
	go run ./cmd/pomogo/main.go

build:
	go build ./cmd/pomogo/main.go

test:
	go test -cover ./...