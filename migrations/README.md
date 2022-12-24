# How to migration

## Databases
There are 1 databases as follows
#### 1. Example Database

## Create migration

1. Example
- Under `example` directory
> $ bash create_migration.sh

and rename file in folder `migrations/example/sql`


## Migrate with docker (Linux)

### Migrate up

- create .env file in migrations folder

```bash
DB_USER=<USER_NAME>
DB_PASSWORD=<PASSWORD>
```

- edit database url in file `migrate_up.sh`
- $ bash migrate_up.sh

## Migrate with cli (Windows, MacOS, or Linux)

### installation

> [https://github.com/golang-migrate/migrate/tree/master/cmd/migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Migrate up

```bash
migrate -path 'path/of/migration_file' \
 -database 'postgres://username:password@ip_address:example_db/db_name?sslmode=disable' \
 -verbose up
```