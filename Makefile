run-all: 
	docker-compose up --force-recreate --build -d

test-coverage:
	go test -cover ./...