# Createing a MONGODB database using docker

## Pulling the image from docker hub
```bash
docker pull mongo:latest
'or directly paste the below command' 
docker run -d -p 2717:2707 --name mongodb mongo:latest
```
## Start the mongo container
```bash
docker start  mongodb
```
## Running MONGODB
```bash
docker exec -it mymongo bash
mongosh
```
