package db

import (
	"fmt"
	"log"

	"github.com/hellskater/omniblog/pkg/config"
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB struct to hold the database instance
type DB struct {
	*gorm.DB
}

func Connect() *DB {
	cfg := config.NewConfig()

	db, err := gorm.Open("postgres", "host=localhost user="+cfg.DBUser+" dbname="+cfg.DBName+" sslmode=disable password="+cfg.DBPassword)

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	fmt.Println("Connection to the database established")

	return &DB{db}
}

func (db *DB) Migrate() {

	// Migrate the schema
	db.AutoMigrate(&models.Follower{
		FollowerID: 1,
		FollowedID: 2,
	})
}

func (db *DB) Close() {
	db.DB.Close()
	fmt.Println("Connection to the database closed")
}
