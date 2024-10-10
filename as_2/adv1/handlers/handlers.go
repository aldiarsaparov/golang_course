package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/aldik/crud_go/models"
)

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users2 (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE NOT NULL,
		age INT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	fmt.Println("Table 'users2' created successfully!")
	return nil
}

func InsertUsersWithTransaction(db *sql.DB, users []models.User) error {
	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Вставка данных
	query := `INSERT INTO users2 (name, age) VALUES ($1, $2)`
	for _, user := range users {
		_, err := tx.Exec(query, user.Name, user.Age)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert user %s: %w", user.Name, err)
		}
	}

	// Завершение транзакции
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Transaction completed successfully!")
	return nil
}

func QueryUsers(db *sql.DB, ageFilter int, limit, offset int) ([]models.User, error) {
	var rows *sql.Rows
	var err error

	// Опциональная фильтрация по возрасту
	if ageFilter > 0 {
		query := `SELECT id, name, age FROM users2 WHERE age = $1 LIMIT $2 OFFSET $3`
		rows, err = db.Query(query, ageFilter, limit, offset)
	} else {
		query := `SELECT id, name, age FROM users2 LIMIT $1 OFFSET $2`
		rows, err = db.Query(query, limit, offset)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(db *sql.DB, id int, newName string, newAge int) error {
	query := `UPDATE users2 SET name = $1, age = $2 WHERE id = $3`
	_, err := db.Exec(query, newName, newAge, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	fmt.Printf("User with ID %d updated successfully!\n", id)
	return nil
}

// func deleteUser(db *sql.DB, id int) error {
// 	query := `DELETE FROM users2 WHERE id = $1`
// 	_, err := db.Exec(query, id)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete user: %w", err)
// 	}

// 	fmt.Printf("User with ID %d deleted successfully!\n", id)
// 	return nil
// }
