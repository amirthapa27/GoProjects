```bash
docker run --name cleanarchidb -e MYSQL_ROOT_PASSWORD=mypass -p 3306:3306 -d mysql:latest
mysql -p
```


```bash
docker run --name cleandb -e MYSQL_ROOT_PASSWORD=mypass -p 5432:5432 -d mysql:latest
su postgres
```