# Query Builder in Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

This Go-based Query Builder is a flexible utility for generating SQL queries dynamically. It provides a convenient way to construct SQL SELECT statements with support for various clauses like SELECT, FROM, WHERE, and ORDER BY.

## Features

- **Dynamic Query Generation**: Build SQL queries on the fly based on your application's specific requirements.
- **Support for Multiple Clauses**: Create complex queries with support for SELECT, FROM, WHERE, and ORDER BY clauses.
- **Variable Arguments**: Use variadic function parameters to pass query parameters and arguments.
- **Type-Safe and Readable**: Ensure type safety while building queries with a readable and concise syntax.

## Installation

To use this query builder in your Go project, you can easily install it using `go get`:

```sh
go get github.com/balaji6666/querybuilder
```

## Usage

Here's how to get started with this query builder:

1. Import the package:

    ```go
    import "github.com/balaji6666/querybuilder"
    ```

2. Create a new query builder instance:

    ```go
    builder := querybuilder.NewSQLBuilder()
    ```

3. Chain methods to construct your SQL query:

    ```go
    query, args := builder.
        Select("id", "name").
        From("players").
        Where("age > ?", 25).
        OrderBy("total_wickets DESC").
        Build()
    ```

4. Execute the query with your database or ORM.

5. Handle the results and enjoy flexible SQL query generation.

## Example

```go
package main

import (
    "fmt"
    "github.com/balaji6666/querybuilder"
)

func main() {
    builder := querybuilder.NewSQLBuilder()

    query, args := builder.
        Select("id", "name").
        From("players").
        Where("age > ?", 25).
        OrderBy("total_wickets DESC").
        Build()

    fmt.Println("Generated SQL query:", query)
    fmt.Println("Query arguments:", args)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the need for a flexible and easy-to-use SQL query builder in Go.

## Get in Touch

If you have questions or suggestions, feel free to [contact us](mailto:youremail@example.com) or [open an issue](https://github.com/your-username/querybuilder/issues).
