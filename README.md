# Go Microservice Template ✨

## An Opinionated Foundation for Your Go Microservices on Kubernetes

This repository provides a robust and well-structured template for building microservices with Go, designed for deployment in a Kubernetes environment.

### Key Features and Why You'll Love It

* **Production-Ready Structure:** A clean and organized project layout (`cmd/`, `internal/`, `pkg/`) promotes maintainability and scalability.
* **Configuration Management:** [Viper](https://github.com/spf13/viper) handles configuration from files (YAML, JSON, TOML), environment variables, making your service highly configurable.
* **Structured Logging:** [slog](https://pkg.go.dev/log/slog) provides structured logging with various output formats, and contextual logging.
* **Health Checks:** Built-in support for Kubernetes liveness, readiness, and startup probes, ensuring smooth deployments and service health.
* **HTTP and gRPC Support:** Includes examples for both HTTP (using [Gorilla Mux](https://github.com/gorilla/mux)) and gRPC servers, allowing you to choose the best communication protocol for your needs.  Middleware for CORS, Logging and Panic recovery are included.
* **Multi-Platform Docker Builds:** `Dockerfile` is optimized for multi-stage builds and multi-platform support (linux/amd64, linux/arm64), creating small and efficient container images.
* **Dependency Management:** Uses Go modules for managing dependencies.
* **Graceful Shutdown:** Implements graceful shutdown of servers, ensuring a smooth exit on termination signals.
* **Cobra CLI:** Uses [Cobra](https://github.com/spf13/cobra) to provide a powerful and easy to use command line interface.
* **Makefile:** Includes a Makefile with convenient commands for building, linting, testing, and cleaning your project.
* **GitHub Actions:** CI/CD workflow for building, linting, testing, and building/pushing Docker images.

### Core Components

Here's a quick overview of the main parts of the template:

* `cmd/service/main.go`:  The entry point of your service.  It uses Cobra for command-line argument parsing and orchestrates the service startup and shutdown.
* `internal/`:  This directory contains your application's private code.
    * `config/`:  Handles loading and managing configuration using Viper.
    * `handlers/`:  Contains HTTP and gRPC handler functions.
    * `health/`:  Implements Kubernetes health check endpoints.
    * `log/`:  Provides a wrapper around `slog` for structured logging, configurable output, format, and context.
    * `server/`:  Sets up and manages HTTP and gRPC servers.
    * `service/`:  Contains the core business logic of your service.
* `pkg/`:  This directory is for reusable libraries that you might share across multiple projects.  It's empty by default in this template.
* `Dockerfile`:  Defines how to build a Docker image for your service.
* `go.mod` and `go.sum`:  Go module files for managing dependencies.
* `app.yaml`:  Example YAML configuration file.  Viper supports other formats too!
* `.github/workflows/build.yaml`:  GitHub Actions workflow for CI/CD.
* `Makefile`: Defines useful commands for local development.

### How to Use This Template

1.  **Clone the Repository:**

    ```bash
    git clone [https://github.com/zaibon/go-template.git](https://github.com/zaibon/go-template.git)
    cd go-template
    ```
2.  **Initialize Your Service:**

    * Modify the `go.mod` file to reflect your service's module path:

        ```bash
        go mod edit -module your-service-name
        ```
    * Run go mod tidy
        ```bash
        go mod tidy
        ```
3.  **Configure Your Service:**

    * Create a configuration file (e.g., `app.yaml`) in the root directory.  See the `app.yaml` example for the structure.
    * Set any necessary environment variables.
4.  **Implement Your Logic:**

    * Write your business logic in the `internal/service/service.go` file.
    * Define your HTTP handlers in `internal/handlers/http.go` and gRPC definitions and handlers in `internal/handlers/grpc.go`.
5.  **Build and Run:**

    * Use the Makefile:
        ```bash
        make build
        ```
    * Or directly:
        ```bash
        go build -o ./bin/your-service ./cmd/service/main.go
        ```
    * Run the service:
        ```bash
        ./bin/your-service
        ```
6.  **Containerize with Docker:**

    * Build the Docker image:
        ```bash
        docker build -t your-service-image .
        ```
    * Run the Docker container:

        ```bash
        docker run -p 8080:8080 -p 9090:9090 your-service-image
        ```
7.  **Deploy to Kubernetes:**

    * Create Kubernetes deployment and service manifests.  (This is beyond the scope of this README, but the template is designed to work well in Kubernetes.)

### Configuration

The service uses Viper for configuration.  You can configure it using:

* Configuration files (YAML, JSON, TOML, etc.)
* Environment variables (which override file settings)

See the `internal/config/config.go` file for how configuration is loaded and structured.

### Logging

The service uses `slog` for logging.  The logger is configured in `cmd/service/main.go` using the options pattern defined in `internal/log/log.go`.

### Health Checks

The `/healthz` endpoint provides a liveness probe, and the `/readyz` endpoint provides a readiness probe for Kubernetes.  You can add your own health check logic in the `internal/health` package.

### HTTP Server

The HTTP server uses Gorilla Mux for routing and includes middleware for:

* CORS
* Logging
* Panic recovery

### gRPC Server

The template includes a basic gRPC server setup.  Define your gRPC service definitions in `.proto` files and generate the Go code using `protoc`.

### Makefile

The Makefile provides the following commands:

* `make all`:  Builds, lints, and tests the service.
* `make build`:  Builds the service.
* `make lint`:  Lints the code using golangci-lint.
* `make test`:  Runs the tests.
* `make clean`:  Removes build artifacts.

### GitHub Actions

The template includes a GitHub Actions workflow (`.github/workflows/build.yaml`) that:

* Builds, lints, and tests the service on pushes and pull requests to the `main` and `develop` branches.
* Builds and pushes a Docker image to Docker Hub on pushes to the `main` branch when the build job is successful.

### Contributing

Contributions are welcome!  Feel free to submit issues or pull requests.
