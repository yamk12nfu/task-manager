# task-manager

### サーバーの起動

`ENV=develop go run app/server.go`

### DB の起動

```
cd backend
docker-compose up -d
```

## golang-migrate

```
migrate -path db/migrations -database 'mysql://root:password@tcp(0.0.0.0:3306)/task-manager-admin' up
```
