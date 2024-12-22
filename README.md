# Service Backend

This project is a web service built using the Fiber framework in Go. It provides CRUD operations for managing services, including creating, updating, deleting, and fetching services with pagination and filtering.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/tushargupta7/kong.git
    cd kong
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the database:
    - Ensure you have a running instance of your database.
    - Update the database configuration in the `config` package.

4. Run the application:
    ```sh
    go run main.go
    ```

## Usage

To start the server, run:
```sh
go run main.go
```

The server will start on the configured port (default is `:8080`). You can then use tools like `curl` or Postman to interact with the API.

## API Endpoints

### Login

- **URL:** `/login`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "username": "your_username",
        "password": "your_password"
    }
    ```
- **Responses:**
    - `200 OK`: Login successful, returns a token.
    - `401 Unauthorized`: Invalid credentials.

### Create a Service

- **URL:** `/service`
- **Method:** `POST`
- **Headers:**
    - `Authorization: Bearer <token>`
- **Request Body:**
    ```json
    {
        "name": "Service Name",
        "description": "Service Description"
    }
    ```
- **Responses:**
    - `201 Created`: Service created successfully.
    - `400 Bad Request`: Invalid request payload.

### Update a Service

- **URL:** `/service/:id`
- **Method:** `PUT`
- **Headers:**
    - `Authorization: Bearer <token>`
- **Request Body:**
    ```json
    {
        "name": "Updated Service Name",
        "description": "Updated Service Description"
    }
    ```
- **Responses:**
    - `200 OK`: Service updated successfully.
    - `400 Bad Request`: Invalid service ID or payload.
    - `404 Not Found`: Service not found.

### Delete a Service

- **URL:** `/service/:id`
- **Method:** `DELETE`
- **Headers:**
    - `Authorization: Bearer <token>`
- **Responses:**
    - `200 OK`: Service deleted successfully.
    - `500 Internal Server Error`: Failed to delete service.

### Get service

- **URL:** `/service`
- **Method:** `GET`
- **Headers:**
    - `Authorization: Bearer <token>`
- **Query Parameters:**
    - `search`: Search term for filtering service.
    - `sort_by`: Field to sort by.
    - `order`: Sort order (`asc` or `desc`).
    - `page`: Page number for pagination.
    - `limit`: Number of items per page.
- **Responses:**
    - `200 OK`: List of service.

### Get a Single Service

- **URL:** `/service/:id`
- **Method:** `GET`
- **Headers:**
    - `Authorization: Bearer <token>`
- **Responses:**
    - `200 OK`: Service details.
    - `400 Bad Request`: Invalid service ID.
    - `404 Not Found`: Service not found.

## Error Handling

Errors are handled using a custom `AppError` struct. The `ErrorHandler` middleware captures these errors and returns a structured JSON response.

Example error response:
```json
{
    "error": "Service not found",
    "code": "ErrServiceNotFound",
    "details": {
        "id": "123"
    }
}
```

Postman Collection
You can find the Postman collection for this project [here](https://github.com/tushargupta7/kong/blob/main/postman_collection.json)
