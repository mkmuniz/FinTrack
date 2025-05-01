# Monthly Expenses API

This project is a simple API built in Go for managing monthly expenses. It allows users to record their expenses, categorize them, and generate reports to analyze their spending habits over time.

## Features

- Create, read, and manage expense records.
- Generate reports to identify the highest and lowest expense categories.
- Track monthly spending trends.

## Technologies Used

- Go (Golang)
- MongoDB for database management
- Gorilla Mux for routing

## Project Structure

```
fintrack-api
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── controllers
│   │   └── expense_controller.go # Handles HTTP requests related to expenses
│   ├── models
│   │   └── expense.go          # Defines the Expense struct
│   ├── repositories
│   │   └── expense_repository.go # Interacts with MongoDB
│   ├── routes
│   │   └── routes.go           # Sets up API routes
│   └── services
│       └── expense_service.go   # Contains business logic for expenses
├── config
│   └── database.go             # MongoDB connection configuration
├── go.mod                       # Module definition and dependencies
├── go.sum                       # Dependency checksums
└── README.md                    # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd fintrack-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Configure MongoDB:**
   Update the `config/database.go` file with your MongoDB connection string.

4. **Run the application:**
   ```
   go run cmd/main.go
   ```

## API Usage

- **POST /expenses**: Create a new expense.
- **GET /expenses**: Retrieve all expenses.
- **GET /expenses/report**: Generate a report of expenses.

## License

This project is licensed under the MIT License.