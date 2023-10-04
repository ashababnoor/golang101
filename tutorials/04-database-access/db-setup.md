# Database Setup for `03-database-access` Tutorial

Create Docker container for MySQL
```bash
docker run --name local-mysql -e MYSQL_ROOT_PASSWORD=anypass -v ~/local-mysql-vol:/var/lib/mysql -p 3306:3306 -d mysql/mysql-server:5.7
```

Open terminal for container to store files
```bash
docker exec -it local-mysql sh
cd home

[ -d projects/personal ] || mkdir -p projects/personal
cd projects/personal

[ -d golang101 ] || mkdir golang101
cd golang101

mkdir -p tutorials/04-database-access/sql 
exit
```

Copy `create-tables.sql` file to docker container
```
docker cp /Users/shabab/Documents/Projects/Personal/golang101/tutorials/04-database-access/sql/create-tables.sql local-mysql:/home/projects/personal/golang101/tutorials/04-database-access/sql/create-tables.sql
```

Open mysql from terminal
```bash
docker exec -it local-mysql mysql -u root -p
```

Run the following mysql commands one after the other
```bash
create database recordings;
use recordings;
source /home/projects/personal/golang101/tutorials/04-database-access/sql/create-tables.sql
select * from album;
exit
```

Get the container IP to be able to connect
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' local-mysql
```

