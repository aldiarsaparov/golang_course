package main
// aldiar saparov
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" 
    "log"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "28081932"
    dbname   = "user_service_db"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatalf("Could not establish a connection: %v", err)
    }
    fmt.Println("Successfully connected to the database!")

    createTable(db)
    insertUser(db, "ALDIAR", 21)
    insertUser(db, "Manson", 57)
    queryUsers(db)
}

func createTable(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users1 (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    );`
    
    _, err := db.Exec(query)
    if err != nil {
        log.Fatalf("Error creating table: %v", err)
    }
    fmt.Println("Table created successfully!")
}

func insertUser(db *sql.DB, name string, age int) {
    query := `
    INSERT INTO users1 (name, age) VALUES ($1, $2);`
    
    _, err := db.Exec(query, name, age)
    if err != nil {
        log.Fatalf("Error inserting user: %v", err)
    }
    fmt.Printf("User %s inserted successfully!\n", name)
}

func queryUsers(db *sql.DB) {
    query := `SELECT id, name, age FROM users1;`
    
    rows, err := db.Query(query)
    if err != nil {
        log.Fatalf("Error querying users1: %v", err)
    }
    defer rows.Close()

    fmt.Println("Users:")
    for rows.Next() {
        var id int
        var name string
        var age int
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            log.Fatalf("Error scanning row: %v", err)
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }
}
