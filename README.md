# GoLang

A repository created for practicing GoLang from beginners to pro.

---

## Overview

This repository is dedicated to learning, practicing, and mastering the Go programming language (Golang). It provides a structured collection of sample programs, real-world use cases, advanced topics, and integration with popular Go tools and frameworks. The codebase is suitable for both beginners and advanced learners looking to strengthen their Go skills through hands-on coding and practical examples.

---

## Table of Contents

- [Project Structure](#project-structure)
- [Technologies & Packages](#technologies--packages)
- [Getting Started](#getting-started)
- [Sample Code](#sample-code)
- [Redis & Docker Integration](#redis--docker-integration)
- [Contributing](#contributing)
- [License](#license)

---

## Project Structure

```
.
├── Example_topics/      # Example Go programs covering language features
├── Tour_of_Go/          # Exercises and code from the official 'A Tour of Go'
├── gfg_golang/          # Advanced topics, generics, web frameworks (gin, gorm, etc.)
├── redis/               # Go-Redis usage and Docker integration
├── go-jwt-master/       # JWT authentication microservice with Docker support
├── tools.txt            # List of major packages and tools used
├── README.md            # Project documentation
```

---

## Technologies & Packages

- **Go (Golang)**
- **Gin** - HTTP web framework
- **Gorm** - ORM library for Go
- **Protobuf** - Protocol Buffers serialization
- **CompileDaemon** - Live reloading for Go apps
- **Redis** - Caching and storage, with Go SDK integration
- **Docker** - Containerization for development and deployment
- **Table Plus** - Database management tool (external utility)
- Additional Go modules: `golang.org/x/net`, `google.golang.org/protobuf`, `google.golang.org/grpc`, and more.

See `tools.txt` for the complete curated list of Go packages and applications.

---

## Getting Started

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/ayushsisodia7/GoLang.git
   cd GoLang
   ```

2. **Explore Example Programs:**
   - Browse `Example_topics/` and `Tour_of_Go/` for beginner to intermediate Go exercises.
   - See `gfg_golang/` for advanced patterns, generics, and using frameworks like Gin.

3. **Run Redis Examples:**
   - Start Docker and Redis as described in `redis/readme.txt`.
   - Run the Go example: `go run redis/main.go`

4. **Try the JWT Auth Microservice:**
   - Build and run the Docker container in `go-jwt-master/` as per the included `Dockerfile`.

---

## Sample Code

### Generics in Go

```go
func add[T ~int](x T, y T) T {
    return x + y
}
```

### Redis Integration Example (see `redis/main.go`)

```go
client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})
// Set and get a serialized object in Redis
err = client.Set(context.Background(), UniqueKey, JsonString, 0).Err()
val, err := client.Get(context.Background(), UniqueKey).Result()
```

---

## Redis & Docker Integration

The `redis/` directory provides step-by-step instructions (`redis/readme.txt`) for setting up a Redis instance locally using Docker, allowing you to experiment with Go and Redis integration in a real environment.

**Example commands:**
```sh
colima start
docker run --name redis-test-instance -p 6379:6379 -d redis
docker ps
docker rm redis-test-instance
colima status
```

---

## Contributing

Contributions are welcome! You can add new Go programs, improve documentation, or suggest new tools and integrations. Please open a pull request with your enhancements.

---

## License

This project is for educational and practice purposes. For licensing, refer to the repository or contact the author.
