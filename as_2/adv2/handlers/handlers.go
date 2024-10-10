package handlers

import (
	"fmt"

	"gorm.io/gorm"
	"github.com/aldik/crud_go/models"
)

func InsertUserWithProfile(db *gorm.DB, user models.User) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to insert user and profile: %w", err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("User and Profile inserted successfully!")
	return nil
}

func GetUsersWithProfiles(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Preload("Profile").Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users with profiles: %w", err)
	}

	fmt.Printf("Retrieved %d users with profiles.\n", len(users))
	return users, nil
}

func UpdateUserProfile(db *gorm.DB, userID uint, newName string, newAge int, newBio string, newProfilePic string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("id = ?", userID).Updates(models.User{Name: newName, Age: newAge}).Error; err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		if err := tx.Model(&models.Profile{}).Where("user_id = ?", userID).Updates(models.Profile{Bio: newBio, ProfilePictureURL: newProfilePic}).Error; err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("User and Profile updated successfully!")
	return nil
}

func DeleteUserWithProfile(db *gorm.DB, userID uint) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.User{}, userID).Error; err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("User and associated Profile deleted successfully!")
	return nil
}
