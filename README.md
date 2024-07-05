# gRPC User Service

This is a simple gRPC service for managing user details.

## Building and Running

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/grpc-user-service.git
   ```
2. Build the project:
   ```sh
   cd grpc-user-service
   go build -o grpc-user-service
   ```
3. Run the server:
   ```sh
   ./grpc-user-service
   ```

## Using Docker

1. Build the Docker image:
   ```sh
   docker build -t grpc-user-service .
   ```
2. Run the Docker container:
   ```sh
   docker run -p 50051:50051 grpc-user-service
   ```

## gRPC Endpoints

- `GetUser`: Fetch user details by ID.
- `GetUsers`: Fetch multiple user details by IDs.
- `SearchUsers`: Search user details based on specific criteria.

## Configuration

No additional configuration is needed.
