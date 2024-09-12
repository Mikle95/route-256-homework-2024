run-all: 
	docker-compose up --force-recreate --build -d

test-coverage:
	go test ./cart/.../service/... ./cart/.../repository/... -coverprofile cover.out.tmp
	type cover.out.tmp | findstr /v "_mock.go" > cover.out
	go tool cover -func cover.out

check-linters:
	golangci-lint run