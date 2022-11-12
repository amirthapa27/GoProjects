# API using MYSQL
This is a project to demonstrate usage of orm....

### Command to start mysql DB
```bash
sudo docker run --name amir -e MYSQL_ROOT_PASSWORD=mypass -p 3306:3306 -d mysql:latest
```

### Run the container
```bash
sudo docker exec -it amir bash
```

### Run MYSQL
```bash 
mysql -u root -p
```

### Create a new databse
```bash
CREATE DATABASE simplerest;
```

### Run main.go

