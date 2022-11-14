# API using PostgreSQL

## Command to start PostgreSql
```bash
docker run --name postgres_db -e POSTGRES_PASSWORD=mypass -p 5432:5432 -d postgres:latest
```
### Run the container 
```bash
docker exec -it postgres_db /bin/bash
```
## Run PostgreSQL 
```bash
su postgres
psql
```
## Create a database is PostgreSQL
```bash
CREATE DATABASE stocksdb;
```
 ## Create a .env file and add 
 ```bash
  POSTGRES_URL = "postgres://root:mypass@localhost:5432/stocksdb"
  ```