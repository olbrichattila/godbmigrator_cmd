migrate:
	go run . migrate
rollback:
	go run . rollback
refresh:
	go run . refresh
install:
	go build -o ./build/migrator .
switch-sqlite:
	cp .env.sqlite.example .env
switch-mysql:
	cp .env.mysql.example .env
switch-pgsql:
	cp .env.pgsql.example .env
