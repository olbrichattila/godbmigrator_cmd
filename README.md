# Golang Database migrator

## Create migration SQL files

If you would like to build the migration into your application, please see:
https://github.com/olbrichattila/godbmigrator/


## What is the provider?
Currently it supports two type of migration provider, json and database.
This is the way the migrator knows which migration was executed and when.

If the json provider is used, then a json file will be saved next to the migration files:
```./migrations/migrations.json```

If the db provider is user, then a migrations table will be created in the same database where you are migrating to.

## Migration file structure
Follow the structure:
[id]-migrate-[custom-content].sql

The files will be processed in ascending order, therefore it is important to create an id as follows:
For example:
```
0001-migrate.sql
0001-rollback.sql
0002-migrate.sql
0002-rollback.sql
0003-migrate.sql
0003-rollback.sql
0004-migrate.sql
0004-rollback.sql
0005-migrate-new.sql
0005-rollback-new.sql
0006-migrate-new.sql
0006-rollback-new.sql
```

## Command line usage:
Migrate:
```go run cmd/cmd.go migrate```

Rollback:
```go run cmd/cmd.go rollback```

Adding new migratio and rollack file:
``````go run cmd/cmd.go add <your custom message>``````
Note: the custom message is not mandatory, in that case the file will be a standard format, like date_time-migration.sql

### Migrate or rollback specified amount of migrations (like 2)
Migrate:
```go run cmd/cmd.go migrate 2```

Rollback:
```go run cmd/cmd.go rollback 2```

### When building the application.
```make install```
The build folder will contain the migrator executable.

Usage is the same but using the application:

```
migrator migrate
migrator rollback

migrator migrate 2
migrator rollback 2
```

The number of rollbacks and migrates are not mandatory.
If it is set, for rollbacks it only apply for the last rollback batch

## .env settings
Create a .env file into your root directory
Examples:

### sqlite
```
DB_CONNECTION=sqlite
DB_DATABASE=./data/database.sqlite
```

### MySql
```
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=migrator
DB_USERNAME=root
DB_PASSWORD=password
```

### Postgres
```
DB_CONNECTION=pgsql
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres
```

Note: Postres currently supports only sslmod disable, others to come:
- disable
- require
- verify-ca
- verify-full
- prefer
- allow

## Setting migration path
The path by default is ./migrations
This can be overwritten by adding the followin variable to your .env file
```
MIGRATOR_MIGRATION_PATH=./migrations/custom_path
```

## Setting the migration provider in .env
It is possible to set the migration provider (see above, saves to database or json)
Possible values are:
```
MIGRATOR_MIGRATION_PROVIDER=json
MIGRATOR_MIGRATION_PROVIDER=db
```
If not set, it defaults to db.

## Available make targets:
```
mage migrate
make rollback
make install
```
## Switch .env file with boilerplate setup for the followin database connections
```
switch-sqlite:
switch-mysql:
switch-pgsql:
```
