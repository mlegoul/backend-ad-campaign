# backend-ad-campaign

A backend API in Go for managing advertising campaigns. Features include campaign creation, updates, search, and deletion. Active campaign filtering is also supported.

## Project Structure

The project follows a clean architecture approach with the following structure :

```
backend-ad-campaign/
├── cmd/
│   └── rest/
│       └── main.go
├── config/
│   └── .env
├── internal/
│   ├── adapters/
│   │   ├── api/
│   │   │   ├── rest.go
│   │   │   └── test/
│   │   │       └── rest_test.go
│   │   └── repository/
│   │       └── postgres.go
│   ├── core/
│   │   └── campaign.go
│   ├── ports/
│   │   └── repository.go
├── mocks/
│   └── mock_repository.go
├── go.mod
├── LICENSE
├── README.md
├── docker-compose.yml
└── Dockerfile
```

## Future Improvements
1: Implement Unit Tests (>80% Coverage)

2: Add Swagger Documentation

3: Integrate Functional Programming

4: Introduce gRPC

5: Connect to a Front (optional)