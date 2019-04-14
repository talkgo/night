# Gin Example

**[Endpoints](ENDPOINTS.md)**

## How to Start

- `make test`
- `make pg-up`
- `make run`
- `make pg-down`

## Development Flow

```
1. Modfiy interfaces
2. Modfiy services
3. Modfiy repositories
4. Modfiy handlers
5. Modfiy routes
```

## Test Flow

```
1. Modfiy mocks
2. Modfiy handler_test
```

**Note**

1. Handlers only interact with Services
2. Services interact with Repository (DB)
