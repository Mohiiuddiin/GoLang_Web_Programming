package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

// Replace with your own connection parameters
var server = "MOHIUDDIN\\SQLEXPRESS" //"202.84.32.123"//
var port = 1433
var user = "sa"
var password = "#Abc123456"

var db *sql.DB

func main() {
    var err error

    // Create connection string
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
        server, user, password, port)

    // Create connection pool
    db, err = sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Error creating connection pool: " + err.Error())
    }
    log.Printf("Connected!\n")

	// createId, err := CreateEmployee("Jake", "United States")
    // fmt.Printf("Inserted ID: %d successfully.\n", createId)

    // Read employees
    count, err := ReadEmployees()
    fmt.Printf("Read %d rows successfully.\n", count)

    // Update from database
    // updateId, err := UpdateEmployee("Jake", "Poland")
    // fmt.Printf("Updated row with ID: %d successfully.\n", updateId)

    // // Delete from database
    // rows, err := DeleteEmployee("Jake")
    // fmt.Printf("Deleted %d rows successfully.\n", rows)
    // // Close the database connection pool after program executes

	if err != nil {
		fmt.Println(err)
	}
	
    defer db.Close()

    // SelectVersion()
}

// Gets and prints SQL Server version
func SelectVersion(){
    // Use background context
    ctx := context.Background()

    // Ping database to see if it's still alive.
    // Important for handling network issues and long queries.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    var result string

    // Run query and scan for result
    err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
    if err != nil {
        log.Fatal("Scan failed:", err.Error())
    }
    fmt.Printf("%s\n", result)
}

func CreateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()
    var err error

    if db == nil {
        log.Fatal("What?")
    }

    // Check if database is alive.
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("INSERT INTO SampleDB.TestSchema.Employees (Name, Location) VALUES ('%s','%s');",
        name, location)

		fmt.Println(tsql)

    // Execute non-query
    result, err := db.ExecContext(ctx, tsql)
    if err != nil {
        log.Fatal("Error inserting new row: " + err.Error())
        return -1, err
    }

    return result.LastInsertId()
}

func ReadEmployees() (int, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("select ID,UserName,Passward from dbCID.dbo.SoftUser;")

    // Execute query
    rows, err := db.QueryContext(ctx, tsql)
    if err != nil {
        log.Fatal("Error reading rows: " + err.Error())
        return -1, err
    }

    defer rows.Close()

    var count int = 0

    // Iterate through the result set.
    for rows.Next() {
        var UserName, Passward string
        var ID int

        // Get values from row.
        //err := rows.Scan(&id, &name, &location)
		err := rows.Scan(&ID,&UserName, &Passward)
        if err != nil {
            log.Fatal("Error reading rows: " + err.Error())
            return -1, err
        }

        fmt.Printf("ID: %d, Name: %s, Password: %s\n", ID, UserName, Passward)
        count++
    }

    return count, nil
}

// Update an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("UPDATE SampleDB.TestSchema.Employees SET Location = @Location WHERE Name= @Name")

    // Execute non-query with named parameters
    result, err := db.ExecContext(
        ctx,
        tsql,
        sql.Named("Location", location),
        sql.Named("Name", name))
    if err != nil {
        log.Fatal("Error updating row: " + err.Error())
        return -1, err
    }

    return result.LastInsertId()
}

// Delete an employee from database
func DeleteEmployee(name string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("DELETE FROM SampleDB.TestSchema.Employees WHERE Name=@Name;")

    // Execute non-query with named parameters
    result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
    if err != nil {
        fmt.Println("Error deleting row: " + err.Error())
        return -1, err
    }

    return result.RowsAffected()
}