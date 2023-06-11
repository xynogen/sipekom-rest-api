## Setup Database
```
docker run -d --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD='yes' -e MYSQL_ROOT_HOST=172.17.0.1 -e MYSQL_DATABASE=sipekom -p 3306:3306 mysql:latest
```

## Compile Go
1. ```go mod tidy```
2. ```go mod vendor```
3. ```go run ./main.go```