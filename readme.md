# DDD with UberFx

This repository contains the source code for "My Awesome Application." It is a Go application structured using the DDD (Domain-Driven Design) approach. The application provides functionality related to managing users.

## Project Structure
```
.
├── config
│   └── config.go
├── domain
│   └── user.go
├── go.mod
├── go.sum
├── handler
│   └── user_handler.go
├── main.go
├── readme.md
├── repository
│   └── user_repository.go
├── service
│   └── user_service.go
└── uberfx
    ├── invoker.go
    └── provider.go

```

- **config:** Contains the configuration package responsible for loading the application configuration.
- **domain:** Holds the domain models, in this case, the `User` model.
- **handler:** Includes the HTTP handler functions for user-related endpoints.
- **repository:** Contains the user repository interface and its implementation.
- **service:** Provides the user service implementation.
- **uberfx:** Includes the UberFx invoker and provider functions for dependency injection.

## Getting Started

To run the application locally, follow these steps:

1. Clone the repository:

2. Install the required dependencies:

3. Set up the application configuration:

4. Build and run the application

5. The application will start running on `http://localhost:8080`. You can now use the API endpoints to manage users.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License

This project is licensed under the [Self Project License](LICENSE).






