.PHONY: up down neo test mocks cover open print clearTestCache

mocks: ## generates interface mocks
	go generate ./...

cover: clearTestCache ## Runs unit tests and creates cover.out file
	go test -v ./... -coverprofile cover.out

open: ## opens coverage report in browser
	go tool cover -html cover.out

print: ## displays coverage percent per func
	go tool cover -func cover.out

clearTestCache:
	go clean -testcache

up: neo ## Start docker containers
	docker-compose up -d

neo: ## Start Neo4j container
	docker run \
    --name neo4j_deputy \
    -p7474:7474 -p7687:7687 \
	-d \
	--rm \
    -v /${HOME}/neo4j/data:/data \
    -v /${HOME}/neo4j/logs:/logs \
    -v /${HOME}/neo4j/import:/var/lib/neo4j/import \
    -v /${HOME}/neo4j/plugins:/plugins \
    --env NEO4J_AUTH=neo4j/test \
	--env NEO4J_apoc_export_file_enabled=true \
    --env NEO4J_apoc_import_file_enabled=true \
    --env NEO4J_apoc_import_file_use__neo4j__config=true \
    --env NEO4JLABS_PLUGINS=\[\"apoc\"\] \
    neo4j:3.5.13

neo_stop:
	docker stop neo4j_deputy

down: neo_stop ## Stop running docker containers
	docker-compose down

test: clearTestCache ## Run unit tests
	go test -v ./...

help: ## Display available commands	
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'