# Create a PostgreSQL database

## Run this command to create a database

```bash
docker run --name menudb -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=mypass -d postgres:latest
```
## Start the databse
```bash 
docker start menudb
```