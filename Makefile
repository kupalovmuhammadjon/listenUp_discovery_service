CURRENT_DIR := $(shell pwd)
DATABASE_URL="postgres://postgres:root@localhost:5432/listenup_discovery_service?sslmode=disable"

run:
	@go run cmd/main.go

gen-proto:
	@./scripts/gen-proto.sh $(CURRENT_DIR)

tidy:
	@go mod tidy
	@go mod vendor

mig-create:
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
	fi; \
	migrate create -ext sql -dir migrations -seq $$name

mig-up:
	@migrate -database "$(DATABASE_URL)" -path migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path migrations down

mig-force:
	@if [ -z "$(version)" ]; then \
		read -p "Enter migration version: " version; \
	fi; \
	migrate -database "$(DATABASE_URL)" -path migrations force $$version

permission:
	@chmod +x scripts/gen-proto.sh
