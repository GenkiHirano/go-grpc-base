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

# TODO: DB関連のコマンド追加する
# setup-db:
# 	@make migrate-down
# 	@make migrate
# 	@make seeder

create-table:
	docker-compose exec sample_app go run internal/database/migrate/create-table/main.go

drop-table:
	docker-compose exec sample_app go run internal/database/migrate/drop-table/main.go

reset-table:
	docker-compose exec sample_app go run internal/database/migrate/reset-table/main.go

proto:
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