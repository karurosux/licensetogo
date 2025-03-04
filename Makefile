export PNAME := licensetogo

run:
	air

build-local:
	make bind-potato \
	&& rm -rf ./dist \
	&& mkdir -p ./dist \
	&& mkdir -p ./dist/client \
	&& cd ./client/ \
	&& npm run build \
	&& cp -r ./dist/* ../dist/client/ \
	&& cd .. \
	&& go build -o ./dist/main ./main.go

build:
	docker build --no-cache -t $$PNAME .

install:
	make build \
	&& docker-compose up -d

gen-types:
	npx pocketbase-typegen@1.2.1 --db $$HOME/$$PNAME/pb_data/data.db --out ./client/src/lib/models/generated/pb-models.ts

migrations-sync:
	go run main.go migrate history-sync --dir $$HOME/$$PNAME/pb_data/

migrate:
	go run main.go migrate collections --dir $$HOME/$$PNAME/pb_data/

create-migration:
	@if [ "$(name)" = "" ]; then \
		echo "Error: Migration name is required. Usage: make create-migration name=your_migration_name"; \
		exit 1; \
	fi
	go run main.go migrate create --dir $$HOME/$$PNAME/pb_data/ "$(name)"

configure:
	docker compose up -d
