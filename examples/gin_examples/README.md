# Gin Example

## Testing

```
make test
```

## Running

```
make run
```

## CURL

Register

```
curl -XPOST -d '{ "name": "heisenberg", "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/register'
```

Login

```
curl -XPOST -d '{ "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/login'
```

## Development Flow

```
1. Add method to interface in user.go
2. Add method to userservice.go
3. Add method to userRepository.go
4. Add method to userHandler.go
5. Add method to routes.go
```

**Note**

1. Handlers only interact with Services
2. Services interact with Repository (DB)
