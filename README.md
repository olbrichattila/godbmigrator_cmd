# Golang Database migrator

## Create migration SQL files

If you would like to build the migration into your application, please see:
https://github.com/olbrichattila/godbmigrator/


## What is the provider?
Currently it supports two type of migration provider, json and database.
This is the way the migrator knows which migration was executed and when.

If the json provider is used, then a json file will be saved next to the migration files:
```
./migrations/migrations.json
```

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
```
go run . migrate
```

Rollback:
```
go run . rollback
```

Report:
```
go run . report
```

Adding new migratio and rollack file:
```
go run . add <your custom message>
```
Note: the custom message is not mandatory, in that case the file will be a standard format, like date_time-migration.sql

### Migrate or rollback specified amount of migrations (like 2)
Migrate:
```
go run . migrate 2
```

Rollback:
```
go run . rollback 2
```
Refresh
(Refresh is when all applied migration is rolled back and migrated up from scratch)
```
go run . refresh
```
Here if the count parameter supplied will be ignored


### When building the application.
```
make install
```
The build folder will contain the migrator executable.

Usage is the same but using the application:

```
migrator migrate
migrator rollback

migrator migrate 2
migrator rollback 2

migrator refresh
migrator report
```

The number of rollbacks and migrates are not mandatory.
If it is set, for rollbacks it only apply for the last rollback batch

## .env settings

If the .env does not exists, the applicaion will read the operating system environment variables.
If the .env file exists and the operating system variables are also set, the operating system variables are taking priority

Example setting variables in linux, command line:
```
export DB_CONNECTION=sqlite
export DB_DATABASE=./data/database.sqlite
```

Unset the variables can be done:
```
unset DB_CONNECTION
unset DB_DATABASE
```

### Create a .env file into your root directory
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

# non mandatory, it defaults to disable
# possible values are: disable, require, verify-ca, verify-full, prefer, allow (depending on your setup)
DB_SSLMODE=disable
```

### Firebird / Interbase
```
DB_CONNECTION=firebird
DB_HOST=127.0.0.1
DB_PORT=3050
DB_DATABASE=/opt/firebird/examples/empbuild/employee.fdb
DB_USERNAME=SYSDBA
DB_PASSWORD=masterkey

MIGRATOR_MIGRATION_PATH=./migrations/firebird
MIGRATOR_MIGRATION_PROVIDER=db
```

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
make refresh
make install
make report
```
## Switch .env file with boilerplate setup for the followin database connections
```
make switch-sqlite
make switch-mysql
make switch-pgsql
make switch-firebird
```

## Test locally
### 1. Docker setup:

```
# Create your docker containers with docker-compose, MySql, FireBird and Postgresql images will be creted exposing the default ports, Change it if required.
cd docker
docker-compose up -d
```

Switch to your testable database with the above make switc-<dbengine> command.

Check your .env file for migration path:

```
MIGRATOR_MIGRATION_PATH=./migrations/new
```

and create a folder if not exists:
```
mkdir -p ./migrations/new
```

Add your migrations:
```
go run . add <optonally a file name suffix>
```

Fill in your migration and rollback file you created, then try migrate, rollback, (with number parameters) and report as explained above

