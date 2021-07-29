DB_USER=root
DB_PASSWORD=secret
DB_NAME=book_management
DDL_SCRIPT=./pkg/db/db_schema.sql

db: run-mysql-docker load-tables

.PHONY: all db

run-mysql-docker:
	docker run --rm -d --name mysql-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD="${DB_PASSWORD}" mysql:8.0.22

load-tables: 
	mysql -u${DB_USER} -p${DB_PASSWORD} -h 127.0.0.1 < ${DDL_SCRIPT}