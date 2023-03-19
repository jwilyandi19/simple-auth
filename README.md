# Simple Auth

## How to Run

Run Database. This will create docker container for mysql and create new database (auth_db)
```
make run-db
```

Migrate all tables to be created in DB, along with seeding some data.
```
make migrate-up
```

Run program
```
make start
```