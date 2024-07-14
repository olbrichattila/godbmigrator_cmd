migrate:
	go run ./cmd/migrator/ migrate
rollback:
	go run ./cmd/migrator/ rollback
refresh:
	go run ./cmd/migrator/ refresh
report:
	go run ./cmd/migrator/ report
run-test:
	go test -v ./...
install:
	go build -o ./build/migrator .
switch-sqlite:
	cp .env.sqlite.example .env
switch-mysql:
	cp .env.mysql.example .env
switch-pgsql:
	cp .env.pgsql.example .env
switch-firebird:
	cp .env.firebird.example .env
lint:
	gocritic check ./...
	revive ./...
	golint ./...
	goconst ./...
	golangci-lint run
	go vet ./...
	staticcheck ./...