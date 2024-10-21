package project

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/denisenkom/go-mssqldb"
)

// User represents a user in the database.
type User struct {
    Id   int
    Name string
    Age  int
}

// getDBConnection establishes a connection to the SQL Server database.
func getDBConnection() (*sql.DB, error) {
    connString := "sqlserver://username:password@localhost:1433?database=yourdb"
    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        return nil, err
    }
    return db, nil
}

// createUser inserts a new user into the database.
func createUser(db *sql.DB, name string, age int) error {
    query := "INSERT INTO Users (Name, Age) VALUES (@Name, @Age)"
    _, err := db.Exec(query, sql.Named("Name", name), sql.Named("Age", age))
    return err
}

// getUsers retrieves all users from the database.
func getUsers(db *sql.DB) ([]User, error) {
    query := "SELECT Id, Name, Age FROM Users"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// updateUser modifies an existing user in the database.
func updateUser(db *sql.DB, id int, name string, age int) error {
    query := "UPDATE Users SET Name = @Name, Age = @Age WHERE Id = @Id"
    _, err := db.Exec(query, sql.Named("Name", name), sql.Named("Age", age), sql.Named("Id", id))
    return err
}

// deleteUser removes a user from the database.
func deleteUser(db *sql.DB, id int) error {
    query := "DELETE FROM Users WHERE Id = @Id"
    _, err := db.Exec(query, sql.Named("Id", id))
    return err
}

// main function demonstrating the CRUD operations.
func main() {
    db, err := getDBConnection()
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    // Create a user
    if err := createUser(db, "Alice", 30); err != nil {
        log.Fatal("Error creating user:", err)
    }

    // Read users
    users, err := getUsers(db)
    if err != nil {
        log.Fatal("Error fetching users:", err)
    }
    fmt.Println("Users:", users)

    // Update a user
    if err := updateUser(db, 1, "Alice Updated", 31); err != nil {
        log.Fatal("Error updating user:", err)
    }

    // Delete a user
    if err := deleteUser(db, 1); err != nil {
        log.Fatal("Error deleting user:", err)
    }
}