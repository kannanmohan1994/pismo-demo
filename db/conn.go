package database

import (
	"fmt"
	"xm/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func PrepareDatabase() (*gorm.DB, error) {
	config := config.GetConfig()

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode = disable password= %s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresDB,
		config.PostgresPassword,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.PostgresSchema + ".",
			SingularTable: false,
		},
	})
	if err != nil {
		fmt.Println("error connecting database - %w", err)
		return nil, err
	}
	fmt.Println("db connection established!")
	return db, nil
}

// func InitialMigration(db *gorm.DB) {
// 	db.AutoMigrate(&model.Customer{}, &model.Order{}, &model.OrderItem{}, &model.Product{})
// }
