# sandbox

HSI Sandbox #3 Golang track.

## How to

Run dev database using `mysql:8.0` image:

```bash
docker run -d --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql:8.0
```

Access the database container:

```bash
docker exec -it mysqldb sh
# or
docker exec -it mysqldb mysql -u root -ppassword
```

Build the app:

```bash
go build .
```

Run the app:

```bash
export DBUSER=root
export DBPASS=password
./sandbox
```
