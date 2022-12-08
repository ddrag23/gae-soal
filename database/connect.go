package database

import (
	"ddrag23/gae-soal/config"
	"ddrag23/gae-soal/model"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func ConnectDB(){
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",config.Config("DB_HOST"), config.Config("DB_USERNAME"), config.Config("DB_NAME"),port)
	DB,err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	  }), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{},&model.Role{})
	fmt.Println("Database Migrated")
}