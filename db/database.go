package db

import (
	"log"
	"microservice_grpc_order/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")
	// dbHost := os.Getenv("DB_HOST")
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	dsn := "root:kl18jda183079@tcp(127.0.0.1:3306)/orders_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return nil, err
	}
	return db, nil
}
