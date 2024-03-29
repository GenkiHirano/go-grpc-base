include .env

.PHONY: proto

init:
	@make up
	@make proto
# @make migrate

rebuild-up:
	@make down
	@make build
	@make up

reup:
	@make down
	@make up

build:
	docker-compose build

up:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

ps:
	docker-compose ps

log:
	docker-compose logs -f sample_app

db-init:
	@make db-down
	@make db-up

db-up:
	docker-compose exec sample_app go run internal/cmd/database/create-table/main.go

db-down:
	docker-compose exec sample_app go run internal/cmd/database/drop-table/main.go

proto-init:
	@make proto-mod-update
	@make proto-lint
	@make proto-gen

proto-mod-update:
	docker-compose run -w /workspace/proto sample_buf mod update

proto-lint:
	docker-compose run --rm sample_buf format -w
	docker-compose run --rm sample_buf lint

proto-gen:
	rm -rf backend/internal/gen
	docker-compose run --rm sample_buf generate --template=buf.gen.yaml --path=proto/sample

go-mod-tidy:
	docker-compose exec sample_app go mod tidy

go-fmt:
	docker-compose exec sample_app go fmt ./...

# TODO: lint実装 (ライブラリ選定含む)

gen-db-schema:
	# must install atlas: https://atlasgo.io/getting-started/
	docker-compose exec sample_app go run internal/cmd/database/create-table/main.go
	atlas migrate diff \
		--dir "file://backend/internal/database" \
		--to "file://backend/internal/schema/000_schema.up.sql" \
		--dev-url "docker://mysql/8/dev" \
		--format '{{ sql . "  " }}'

db-inspect:
	docker-compose exec sample_db atlas schema inspect \
		-u "mysql://sample:sample123@localhost:3306/sample_user" \
		--web