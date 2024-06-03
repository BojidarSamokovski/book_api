
# Book API

This is a simple web API for managing a collection of books, built with Go and Gorilla Mux. The API supports basic CRUD operations: create, read, update, and delete books.

## Project Structure

```
book_api/
├── main.go
├── go.mod
├── router/
│   └── router.go
├── controllers/
│   └── bookController.go
├── models/
│   └── book.go
│   └── book_test.go
└── utils/
    └── responses.go
    └── responses_test.go
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-username/book_api.git
    cd book_api
    ```

2. Initialize the Go module:

    ```sh
    go mod init book_api
    go mod tidy
    ```

3. Install the Gorilla Mux package:

    ```sh
    go get -u github.com/gorilla/mux
    ```

### Running the Server

To start the server, run:

```sh
go run main.go
```

The server will start on port 8080. You should see output indicating the server is running:

```sh
Listening on :8080...
```

### API Endpoints

#### Get All Books

- **Method**: GET
- **URL**: `/books`
- **Description**: Retrieves a list of all books.

#### Get a Single Book

- **Method**: GET
- **URL**: `/books/{id}`
- **Description**: Retrieves a single book by its ID.

#### Create a New Book

- **Method**: POST
- **URL**: `/books`
- **Description**: Creates a new book.
- **Body**: Raw JSON

    ```json
    {
        "title": "New Book",
        "isbn": "111222333",
        "author": "New Author",
        "year": 2023
    }
    ```

#### Update an Existing Book

- **Method**: PUT
- **URL**: `/books/{id}`
- **Description**: Updates an existing book by its ID.
- **Body**: Raw JSON

    ```json
    {
        "title": "Updated Book",
        "isbn": "111222333",
        "author": "Updated Author",
        "year": 2024
    }
    ```

#### Delete a Book

- **Method**: DELETE
- **URL**: `/books/{id}`
- **Description**: Deletes a book by its ID.

### Running Tests

To run all tests, execute the following command in your project root directory:

```sh
go test ./...
```

To run tests for a specific package:

```sh
cd models
go test

cd ../utils
go test
```

### Example `curl` Commands

- **Get all books**:

    ```sh
    curl -X GET http://localhost:8080/books
    ```

- **Get a single book by ID**:

    ```sh
    curl -X GET http://localhost:8080/books/1
    ```

- **Create a new book**:

    ```sh
    curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title":"New Book","isbn":"111222333","author":"New Author","year":2023}'
    ```

- **Update a book by ID**:

    ```sh
    curl -X PUT http://localhost:8080/books/1 -H "Content-Type: application/json" -d '{"title":"Updated Book","isbn":"111222333","author":"Updated Author","year":2024}'
    ```

- **Delete a book by ID**:

    ```sh
    curl -X DELETE http://localhost:8080/books/1
    ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
