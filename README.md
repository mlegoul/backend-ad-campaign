# backend-ad-campaign

A backend API in Go for managing advertising campaigns. Features include campaign creation, updates, search, and deletion. Active campaign filtering is also supported.

## Project Structure

The project follows a clean architecture approach with the following structure :

```
/backend-ad-campaign
├── cmd/
│   └── rest/
│       └── main.go
├── internal/
│   ├── core/
│   │   └── campaign.go
│   ├── ports/
│   │   ├── repository.go
│   │   └── service.go
│   └── adapters/
│       ├── api/
│       │   └── rest.go
│       └── repository/
│           └── postgres.go
├── config/
│   └── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

