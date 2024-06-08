package config

import (
	"github.com/brunobotter/ecommerce-produto/scheamas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeSql() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	dsn := "root:171191@tcp(127.0.0.1:3306)/produto?charset=utf8mb4&parseTime=true&loc=Local"
	//create db and connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql produto error: %v", err)
		return nil, err
	}
	//migrate scheama
	err = db.AutoMigrate(&scheamas.Produto{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
