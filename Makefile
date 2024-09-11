run-all: 
	docker-compose up --force-recreate --build -d

test-coverage:
	go test -cover ./cart/internal/pkg/cart/... ./cart/internal/adapter/product/service