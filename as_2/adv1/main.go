package main
// aldiar saparov
import (
	"fmt"
	"log"
	"github.com/aldik/crud_go/database"
	"github.com/aldik/crud_go/handlers"
	"github.com/aldik/crud_go/models"
)



func main() {
	// Подключение к базе данных
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Создание таблицы
	err = handlers.CreateTable(db)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// Вставка пользователей с транзакцией
	users := []models.User{
		{Name: "Haruki", Age: 75},
		{Name: "Aldik", Age: 25},
		{Name: "Gabdyq", Age: 20},
	}

	err = handlers.InsertUsersWithTransaction(db, users)
	if err != nil {
		log.Fatal("Failed to insert users:", err)
	}

	// Запрос пользователей с фильтрацией и пагинацией
	result, err := handlers.QueryUsers(db, 0, 2, 0)
	if err != nil {
		log.Fatal("Failed to query users:", err)
	}
	fmt.Println("Users:", result)

	err = handlers.UpdateUser(db, 5, "Arthur", 54)
	if err != nil {
		log.Fatal("Failed to update user:", err)
	}

	// Удаление пользователя
	// err = deleteUser(db, 2)
	// if err != nil {
	// 	log.Fatal("Failed to delete user:", err)
	// }
}
