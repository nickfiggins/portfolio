// models/setup.go

package models

import (
	"fmt"
	_ "log"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DynamoDB *dynamo.DB

func ConnectDatabase() {
	DynamoDB = Load_Dynamo()
	err := godotenv.Load("prod.env"); if err != nil {
		fmt.Println(err)
	}
	/*
	TODO: remove !
	
	dsn := os.Getenv("PORTFOLIO_DB_DSN")
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn); if err != nil {
		fmt.Println(err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{PrepareStmt: true})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully connected!")

	DB = gormDB
	*/
}