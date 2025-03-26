# Golang Database Migrator (godbmigrator)  
A simple and flexible database migration tool for Go applications.

## Create Migration SQL Files

## Installation

To install as a CLI tool:
```bash
go install github.com/olbrichattila/godbmigrator_cmd/cmd/migrator@latest
```

If you'd like to integrate the migration into your application, please refer to:
https://github.com/olbrichattila/godbmigrator/


## Migration File Structure
Follow this structure:
[id]-migrate-[custom-content].sql

The files are processed in ascending order, so it's important to prefix them with an ID or timestamp.  
Alternatively, you can use the following command to generate a properly formatted migration file: ```migrator add <optional suffix>```

For example:
```bash
2024-05-27_19_49_38-migrate.sql
2024-05-27_19_49_38-rollback.sql
2024-05-27_19_50_04-migrate.sql
2024-05-27_19_50_04-rollback.sql
```

## Command line usage:
Migrate:
```bash
migrator migrate
```

Rollback:
```bash
migrator rollback
```

Report:
```bash
migrator report
```

Adding new migration and rollback file:
```bash
migrator add <your custom message>
```
Note: the custom message is not mandatory, in that case the file will be a standard format, like date_time-migration.sql


### Running Migrations and Rollbacks  

You can apply or roll back a specific number of migrations by passing a number as an argument.

**Migrate the last 2 migrations:**  
```bash
migrator migrate 2
migrator rollback 2
```


### Refresh
(Refresh is when all applied migration is rolled back and migrated up from scratch)
```bash
migrator refresh
```

Here if the count parameter supplied will be ignored

## Baseline
> Note: this feature is currently in beta
You can create a snapshot of the current database schema and restore it when recreating the database.  
This is also useful for generating a test database from a production database without copying data.
Please note: this supports only SQLite, MySql and PostgreSQL. (firebird support coming later)

Usage:
```bash
migrator save-baseline
migrator restore-baseline
```

### When building the application.
```bash
make install
```
The build folder will contain the migrator executable.

Usage is the same but using the application:

```bash
migrator migrate
migrator rollback
migrator migrate 2
migrator migrate 2 -force
migrator rollback 2
migrator refresh
migrator report
migrator validate
migrator save-baseline
migrator restore-baseline
migrator add <optional migration file suffix>
help (for full detailed help)
```

### Available flags:
```-force``` This flag will skip checksum verification

The number of rollbacks and migrates are not mandatory.
If it is set, for rollbacks it only apply for the last rollback batch
Validate checks if any migration file changed since last applied

## Configuring Database Connections

### **Using Environment Variables**
If no `.env.migrator` file is present, the application will read environment variables from the operating system.

#### **Example (Linux/macOS Terminal)**
```bash
export DB_CONNECTION=sqlite
export DB_DATABASE=./data/database.sqlite
```

Unset the variables can be done:
```bash
unset DB_CONNECTION
unset DB_DATABASE
```

### Create a .env file into your root directory
Examples:
Note: ```TABLE_PREFIX``` is non mandatory, if not set, the migration table prefix will be ```olb``` (example ```olb_migrations```, ```olb_migration_reports```)

### sqlite
```bash
DB_CONNECTION=sqlite  
DB_DATABASE=./data/database.sqlite  
TABLE_PREFIX="my_prefix"  # (Optional: Defaults to "olb_")
```

### MySql
```bash
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=migrator
DB_USERNAME=root
DB_PASSWORD=password
TABLE_PREFIX="my_prefix"
```

### Postgres
```bash
DB_CONNECTION=pgsql
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres
TABLE_PREFIX="my_prefix"

# non mandatory, it defaults to disable
# possible values are: disable, require, verify-ca, verify-full, prefer, allow (depending on your setup)
DB_SSLMODE=disable
```

### Firebird / Interbase
```bash
DB_CONNECTION=firebird
DB_HOST=127.0.0.1
DB_PORT=3050
DB_DATABASE=/opt/firebird/examples/empbuild/employee.fdb
DB_USERNAME=SYSDBA
DB_PASSWORD=masterkey
TABLE_PREFIX="my_prefix"

MIGRATOR_MIGRATION_PATH=./migrations/firebird
```

## Setting migration path
The path by default is ./migrations
This can be overwritten by adding the followin variable to your .env file
```bash
MIGRATOR_MIGRATION_PATH=./migrations/custom_path
```

## HTTP Server Mode
Set the server port. If not defined, it will default to **8080**.
```bash
HTTP_SERVER_PORT=8081
```

## Available make targets:
```bash
make migrate
make rollback
make refresh
make install
make report
```
## Switch .env file with boilerplate setup for the followin database connections
```bash
make switch-sqlite
make switch-mysql
make switch-pgsql
make switch-firebird
```

## Test locally
### 1. Docker setup:

```bash
# Create your docker containers with docker-compose, MySql, FireBird and Postgresql images will be creted exposing the default ports, Change it if required.
cd docker
docker-compose up -d
```

Switch to your testable database with the above make switc-<dbengine> command.

Check your .env file for migration path:

```bash
MIGRATOR_MIGRATION_PATH=./migrations/new
```

and create a folder if not exists:
```bash
mkdir -p ./migrations/new
```

#### Add your migrations:
```bash
migrator add <optionally a file name suffix>
```

Fill in your migration and rollback file you created, then try migrate, rollback, (with number parameters) and report as explained above

---

## HTTP Server
To run the migrator locally, use:
```bash
migrator serve
```

### URL Parameters

The following URL parameters can be used:
- **command** – The migration command (e.g., migrate, rollback).
- **count** (optional) – Number of items to migrate or rollback.
- **force** (optional) – Accepts 0, 1, false, or true. Forces migration even if integrity validation fails.

**Example Requests**
```bash
http://localhost:8081?command=migrate&count=2&force=true
http://localhost:8081?command=add&name=add_users
```

---

## Pre-Built Docker Image
**Pull the Image**
```bash
docker pull aolb/migrator
```
**Run with Environment Variables**
```bash
docker run -e HTTP_SERVER_PORT=8080 -e DB_CONNECTION=sqlite -e DB_DATABASE=./database.sqlite -e MIGRATOR_MIGRATION_PATH=./ -p8080:8080 aolb/migrator
```

**Example docker compose**
```yml
version: "3.8"

services:
  migrator:
    image: aolb/migrator:latest
    build: .
    ports:
      - "8081:8080"
    environment:
      - HTTP_SERVER_PORT=8080
      - DB_CONNECTION=sqlite
      - DB_DATABASE=./database.sqlite
      - MIGRATOR_MIGRATION_PATH=./
    restart: unless-stopped
```

---

## About me:
- Learn more about me on my personal website: https://attilaolbrich.co.uk/menu/my-story
- Check out my latest blog post on my website: https://attilaolbrich.co.uk/blog/1/single
