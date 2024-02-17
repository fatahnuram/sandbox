# sandbox

HSI Sandbox #3 Golang track.

## Level 4, how to

Run dev database using `mysql:8.0` image:

```bash
docker run -d --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql:8.0
```

Access the database container:

```bash
docker exec -it mysqldb mysql -u root -ppassword
```

Create initial database name:

```sql
CREATE DATABASE dbname;
USE dbname;
```

Apply database schema:

```bash
# copy paste content from `hsi-sandbox-lv3.sql` inside `sql` dir.
```

Build the app:

```bash
go build .
```

Export necessary configs:

```bash
export DBUSER=root
export DBPASS=password
export DBNAME=dbname
```

Run the app:

```bash
./sandbox
```
