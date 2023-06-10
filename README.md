## Setup Database
```docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=sipekom -p3306:3306 mysql```

## Compile Go
1. ```go mod tidy```
2. ```go mod vendor```
3. ```go run ./main.go```