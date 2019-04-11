# Gin Example

## Running

```
make pg-up

go get

go run cmd/server/main.go
```

## CURL

```
curl -XPOST -d '{ "name": "heisenberg", "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/register'
```
