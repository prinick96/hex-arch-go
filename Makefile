# Docker
docker-build:
	docker build -t hex-arch-go . 

docker-up: docker-build docker-down-prod
	docker-compose --env-file .env.development up  -d

docker-up-prod: docker-build docker-down
	docker-compose --env-file .env.production up -d

docker-down:
	docker-compose --env-file .env.development down

docker-down-prod:
	docker-compose --env-file .env.production down

# Compilation
compile-win: 
	go build -o bin/hex-arch-go.exe -v

compile:
	go build -o bin/hex-arch-go -v
	
# Heroku
heroku-local:
	cp .env.development bin/.env.development
	heroku local

heroku-run: compile heroku-local

heroku-run-win: compile-win heroku-local

# Testing
unit-test:
	go test ./internal/helpers/tests
	go test ./core/application/...

integration-test:
	go test ./core/infrastructure/...

test: unit-test integration-test