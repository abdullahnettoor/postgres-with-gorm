package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	dbUri := "user=abdullahnettoor dbname=test host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Println("Error occured while connecting to DB", err)
	}

	err = db.AutoMigrate(User{})
	if err != nil {
		fmt.Println("\n\nError migrating user:", err)
	} else {
		fmt.Println("\n\nError is", err)
	}

	user := User{
		FirstName:   "Abdullah",
		LastName:    "Nettoor",
		Email:       "abdullahnttoor@gmail.com",
		DateOfBirth: time.Date(1997, 12, 06, 0, 0, 0, 0, time.UTC),
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if db.Error != nil {
		fmt.Println("Error Creating user:", err)
		return
	}

	email := "abdullahnettoor@gmail.com"
	lname := "Nettoooor"

	db.Model(&user).Where("email = ?", email).Update("last_name", lname)

}

// - DEFAULT
// - VIEW
// - CASE
// - SQL INJECTION
// - NORMALISATION
// - DATA INTEGRITY
// - ROLLBACK
// - UNION
// - COPY TABLE
