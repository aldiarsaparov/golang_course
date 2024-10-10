package main
// aldiar saparov
import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
    Age  int
}

func (User) TableName() string {
    return "users1"
}

func main() {
    dsn := "host=localhost user=postgres password=28081932 dbname=user_service_db port=5432 sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatalf("Error with automigration: %v", err)
    }
    fmt.Println("Table users1 created or already exists.")

    insertUser(db, "al-arbuzi", 1)
    insertUser(db, "murakami", 75)
    insertUser(db, "kafka", 15)
    insertUser(db, "ale", 42)

    queryUsers(db)
}

func insertUser(db *gorm.DB, name string, age int) {
    user := User{Name: name, Age: age}
    result := db.Create(&user)
    if result.Error != nil {
        log.Fatalf("Error inserting user: %v", result.Error)
    }
    fmt.Printf("User %s inserted successfully!\n", name)
}

func queryUsers(db *gorm.DB) {
    var users []User
    result := db.Find(&users)
    if result.Error != nil {
        log.Fatalf("Error querying users: %v", result.Error)
    }

    fmt.Println("Users1 Table:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
    }
}
