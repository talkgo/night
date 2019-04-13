# Endpoints

## Registration

- `/api/register` Register a new user.

```
curl -XPOST -d '{ "name": "heisenberg", "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/register'
```

- `/api/login` Login an existing user.

```
curl -XPOST -d '{ "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/login'
```

- `/api/logout` Logout current user.

```
curl -XPOST 'http://localhost:8080/api/logout'
```

## User

- `/api/v1/me` GET own user (based on SessionToken).

```
curl -v --cookie "sessionID=710ddd62-4f10-4bf3-8f30-325f7f4a297f" 'http://localhost:8080/api/v1/me'
```


- `/api/v1/users/:id` GET users.

```
curl -v --cookie "sessionID=710ddd62-4f10-4bf3-8f30-325f7f4a297f" 'http://localhost:8080/api/v1/users/1'
```
