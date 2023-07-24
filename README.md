# Golang Database migrator

!!Under development...

## Create migration SQL files into the the folder migrations (will be configurable)

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

# Without building the app

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

## Setting the migration provider in .env
It is possible to set the migration provider (see above, saves to database or json)
Possible values are:
```
MIGRATOR_MIGRATION_PROVIDER=json
MIGRATOR_MIGRATION_PROVIDER=db
```
If not set, it defaults to db.

## Adding to your code.

Import the module:

```migrator "github.com/olbrichattila/godbmigrator"```

You need to have a DB connection, and a migration provider.

The migration provider stores the migration status to:
- json
- database
- (others to come)


Currently the command line utility supports only SqLite, the build in solution shoud work, but not tested with oher databases

Coming soon:
- MySql
- Postgresql

## Example migrate: (where the db is your *sql.DB)

```
migrationProvider, err := migrator.NewMigrationProvider("json")
if err != nil {
    panic("Error: " + err.Error())
}

err = migrator.Migrate(db, migrationProvider, count)
if err != nil {
    panic("Error: " + err.Error())
}
```

## Example rollback: (where the db is your *sql.DB)
```
migrationProvider, err := migrator.NewMigrationProvider("json", nil)
if err != nil {
    panic("Error: " + err.Error())
}

err = migrator.Rollback(db, migrationProvider, count)
if err != nil {
    panic("Error: " + err.Error())
}
```

## Migrate With database provider
```
db, err := migrator.NewSqliteStore("./data/database.sqlite")
if err != nil {
    panic("Error: " + err.Error())
}

migrationProvider, err := migrator.NewMigrationProvider("db", db)
if err != nil {
    panic("Error: " + err.Error())
}

err = migrator.Migrate(db, migrationProvider, count)
if err != nil {
    panic("Error: " + err.Error())
}
```

## Rollback With database provider
```
db, err := migrator.NewSqliteStore("./data/database.sqlite")
if err != nil {
    panic("Error: " + err.Error())
}

migrationProvider, err := migrator.NewMigrationProvider("db", db)
if err != nil {
    panic("Error: " + err.Error())
}

err = migrator.Rollback(db, migrationProvider, count)
if err != nil {
    panic("Error: " + err.Error())
}
```

## Example, create new migration file:
(same as ``````go run cmd/cmd.go add <your custom message>``````)
```
err := migrator.AddNewMigrationFiles(secondParameter)
if err != nil {
    panic("Error: " + err.Error())
}
```

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
## Coming soon

.env where you can define the database connection. migration file paths and migration provider type therefore it can be used as a full featured command line migrator.

