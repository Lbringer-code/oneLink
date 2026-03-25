# Backend Project Structure
 
```
backend/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── cleanup/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── handlers.go
│   │   └── router.go
│   ├── config/
│   │   └── config.go
│   ├── service/
│   │   └── service.go
│   ├── repository/
│   │   └── repository.go
│   └── domain/
│       ├── link.go
│       └── bundle.go
├── sql/
│   └── migrations/
├── api/
│   └── openapi.yaml
└── .env
```