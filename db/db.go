package db

import (
	"fmt"
	"hashing-file/config"
	"hashing-file/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDB() *gorm.DB {
	configEnv := config.LoadEnv()
	loadConfig := map[string]string{
		"Username": configEnv.GetString("DB_USERNAME"),
		"Passowrd": configEnv.GetString("DB_PASSWORD"),
		"Host":     configEnv.GetString("DB_HOST"),
		"Port":     configEnv.GetString("DB_PORT"),
		"DB":       configEnv.GetString("DB_NAME"),
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", loadConfig["Username"], loadConfig["Password"], loadConfig["Host"], loadConfig["Port"], loadConfig["DB"])
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	MigrateDB()
	return db
}

func MigrateDB() {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.File{})
}
