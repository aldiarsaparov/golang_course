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
	db, err := database.ConnectGORM()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal("Failed to migrate models:", err)
	}

	user := models.User{
		Name: "Aldiar",
		Age:  21,
		Profile: models.Profile{
			Bio:               "Aspiring student from Kazakhstan",
			ProfilePictureURL: "https://www.starmessagesoftware.com/myfiles/apple-silicon-m1-chip-arm64.jpg",
		},
	}

	err = handlers.InsertUserWithProfile(db, user)
	if err != nil {
		log.Fatal("Failed to insert user and profile:", err)
	}

	users, err := handlers.GetUsersWithProfiles(db)
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}
	fmt.Println("Users with Profiles:", users)

	err = handlers.UpdateUserProfile(db, users[0].ID, "Aldiar Saparov", 210, "old", "https://www.starmessagesoftware.com/myfiles/apple-silicon-m1-chip-arm64.jpg")
	if err != nil {
		log.Fatal("Failed to update user profile:", err)
	}

	// err = deleteUserWithProfile(db, users[0].ID)
	// if err != nil {
	// 	log.Fatal("Failed to delete user:", err)
	// }
}
