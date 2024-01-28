include .env

.PHONY: proto

build:
	docker-compose build

rebuild-up:
	@make down
	@make build
	@make up

init:
	@make up
	@make proto
# @make migrate

# setup-db:
# 	@make migrate-down
# 	@make migrate
# 	@make seeder

up:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

reup:
	@make down
	@make up

ps:
	docker-compose ps

log:
	docker-compose logs -f sample_app

proto:
	rm -rf backend/internal/gen
	docker-compose run --rm sample_buf generate --template=buf.gen.yaml --path=proto/sample

proto-mod-update:
	docker-compose run -w /workspace/proto sample_buf mod update

proto-check:
	docker-compose run --rm sample_buf format -w
	docker-compose run --rm sample_buf lint
