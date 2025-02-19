## My Go Data Generator

This project is designed to generate a large dataset (up to 500GB) in a MySQL database using Go and GORM. It includes functionality for batch inserting data, writing to CSV files, and performing database migrations.

##

运维debezium和flinkcdc的最佳实践
想知道databend的这些方法哪些是增量的那些是全量的

### Project Structure

```
my-go-data-generator
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── db
│   │   ├── connection.go # Database connection logic
│   │   └── migrate.go    # Database migration handling
│   ├── models
│   │   ├── order.go      # Order model definition
│   │   ├── product.go    # Product model definition
│   │   └── user.go       # User model definition
│   ├── generator
│   │   └── generator.go   # Data generation logic
│   └── csv
│       └── writer.go      # CSV writing functionality
├── go.mod                # Go module file
└── README.md             # Project documentation
```

### Setup Instructions

1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**: Clone this repository to your local machine.

   ```
   git clone <repository-url>
   cd my-go-data-generator
   ```

3. **Install Dependencies**: Navigate to the project directory and run the following command to install the necessary dependencies.

   ```
   go mod tidy
   ```

4. **Configure Database**: Update the database connection settings in `internal/db/connection.go` to match your MySQL configuration.

5. **Run Migrations**: Before generating data, run the migrations to set up the database schema.

   ```
   go run cmd/main.go migrate
   ```

6. **Generate Data**: To start generating data, run the following command:

   ```
   go run cmd/main.go generate
   ```

### Usage

- The application will generate data for three tables: `orders`, `products`, and `users`.
- Each table will have meaningful fields and relationships.
- Data will be inserted in batches for efficiency.
- Generated data will also be written to CSV files for easy access and import into MySQL.

### Logging and Progress

The application will log the progress of data generation and insertion, including the number of records generated and the time taken for the operation.

### Future Enhancements

- Implement a feature to periodically insert additional records into the database.
- Add more meaningful fields and relationships to the data models.
- Enhance CSV writing functionality to support different formats.

### License

This project is licensed under the MIT License. See the LICENSE file for more details.