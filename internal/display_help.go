package migrator

import "fmt"

func displayFullHelp() {
	fmt.Printf(`
If the .env.migrator does not exists, the application will read the operating system environment variables.
If the .env.migrator file exists and the operating system variables are also set, the operating system variables are taking priority

Example setting variables in linux, command line:

------------------------------------------
export DB_CONNECTION=sqlite
export DB_DATABASE=./data/database.sqlite
------------------------------------------

Unset the variables can be done:
--------------------
unset DB_CONNECTION
unset DB_DATABASE
------------------

Create a .env file into your root directory
Examples:

sqlite
-------------------------------------
DB_CONNECTION=sqlite
DB_DATABASE=./data/database.sqlite
-------------------------------------

MySql
-------------------------------------
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=migrator
DB_USERNAME=root
DB_PASSWORD=password
-------------------------------------

Postgres
-------------------------------------
DB_CONNECTION=pgsql
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres


# non mandatory, it defaults to disable
# possible values are: disable, require, verify-ca, verify-full, prefer, allow (depending on your setup)
DB_SSLMODE=disable
-------------------------------------

Firebird / Interbase
-------------------------------------
DB_CONNECTION=firebird
DB_HOST=127.0.0.1
DB_PORT=3050
DB_DATABASE=/opt/firebird/examples/empbuild/employee.fdb
DB_USERNAME=SYSDBA
DB_PASSWORD=masterkey

MIGRATOR_MIGRATION_PATH=./migrations/firebird
MIGRATOR_MIGRATION_PROVIDER=db
-------------------------------------

Setting migration path
The path by default is ./migrations
This can be overwritten by adding the following variable to your .env file
-------------------------------------
MIGRATOR_MIGRATION_PATH=./migrations/custom_path
-------------------------------------

Setting the migration provider in .env
It is possible to set the migration provider (see above, saves to database or json)
Possible values are:
-------------------------------------
MIGRATOR_MIGRATION_PROVIDER=json
MIGRATOR_MIGRATION_PROVIDER=db
-------------------------------------
`)
}
