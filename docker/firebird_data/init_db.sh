#!/bin/bash
if [ ! -f /firebird/data/employee.fdb ]; then
  echo "Creating Firebird database..."
  echo "CREATE DATABASE '/firebird/data/employee.fdb';" | /usr/local/firebird/bin/isql -user sysdba -password masterkey
else
  echo "Database already exists."
fi
