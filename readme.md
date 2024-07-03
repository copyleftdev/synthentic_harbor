# Synthetic Harbor Dental Inquiry API

This API synthetically generates patient data for dental inquiries. It provides endpoints to retrieve detailed information about dental inquiries, including plan information, treatment history, and more. This is useful for testing and simulating different scenarios in dental insurance workflows.

## Table of Contents

- [Synthetic Harbor Dental Inquiry API](#synthetic-harbor-dental-inquiry-api)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
    - [API Endpoints](#api-endpoints)
  - [Customization](#customization)
    - [Extending Insurance Providers](#extending-insurance-providers)
    - [Modifying Inquiry Data](#modifying-inquiry-data)
  - [Docker](#docker)
    - [Dockerfile](#dockerfile)
    - [Docker Compose](#docker-compose)

## Getting Started

### Prerequisites

- Go 1.21.4 or later
- Git

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/copyleftdev/synthentic_harbor.git
   cd synthentic_harbor
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Run the application:

   ```sh
   go run main.go
   ```

The application will be available at `http://localhost:8000`.

## Usage

### API Endpoints

- `GET /api/v0/dental/inquiries/:id`
  - Retrieves a dental inquiry by its ID.

  Example:
  ```sh
  curl -X GET "http://localhost:8000/api/v0/dental/inquiries/123e4567-e89b-12d3-a456-426614174000"
  ```

- `GET /api/v0/dental/inquiries`
  - Retrieves a list of dental inquiries (default limit: 10).

  Example:
  ```sh
  curl -X GET "http://localhost:8000/api/v0/dental/inquiries"
  ```

## Customization

### Extending Insurance Providers

To add more insurance providers, update the `insuranceProviders` slice in `handlers/inquiries.go`:

```go
var insuranceProviders = []string{"Aetna", "Cigna", "Humana", "MetLife", "United Healthcare", "Delta Dental", "Blue Cross Blue Shield", "NewProvider1", "NewProvider2"}
```

### Modifying Inquiry Data

The `generateFakeInquiry` function in `handlers/inquiries.go` generates synthetic data for inquiries. To modify the generated data, adjust the fields and logic within this function. For example, to change how the status is generated:

```go
func randomStatus() string {
    statuses := []string{"SCHEDULED", "IN_PROGRESS", "SUCCESS", "UNSUCCESSFUL", "UNKNOWN"}
    return statuses[rand.Intn(len(statuses))]
}
```

## Docker

### Dockerfile

A `Dockerfile` is provided to containerize the application.

```dockerfile
# Start from the official Golang base image for version 1.21.4
FROM golang:1.21.4-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
```

### Docker Compose

A `docker-compose.yml` file is provided to run the application using Docker Compose.

```yaml
version: '3.8'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8000:8000"
```

To run the application using Docker Compose:

```sh
docker-compose up --build
```

