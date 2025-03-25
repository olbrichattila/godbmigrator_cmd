migrate:
	go run ./cmd/migrator/ migrate
rollback:
	go run ./cmd/migrator/ rollback
refresh:
	go run ./cmd/migrator/ refresh
report:
	go run ./cmd/migrator/ report
test:
	go test -v ./...
install:
	go build -o ./build/migrator .
switch-sqlite:
	cp .env.sqlite.example .env.migrator
switch-mysql:
	cp .env.mysql.example .env.migrator
switch-pgsql:
	cp .env.pgsql.example .env.migrator
switch-firebird:
	cp .env.firebird.example .env.migrator
docker-build:
	docker build -t migrator .
docker-run:
	docker run -d --name migrator -p 8081:8080 migrator
lint:
	# gocritic check ./...
	revive ./...
	golint ./...
	goconst ./...
	# golangci-lint run
	go vet ./...
	staticcheck ./...
	