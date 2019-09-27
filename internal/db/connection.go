package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/server-may-cry/go-workers-test/internal/entity"
	"log"
)

type Connection struct {
	db *gorm.DB
}

type ConnectionConfiguration struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func OpenConnection(conf ConnectionConfiguration) *Connection {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		//"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	))
	if err != nil {
		log.Panicf("failed to connect database %+v", err)
	}
	db.LogMode(true)

	db.AutoMigrate(&entity.MyEntry{})

	return &Connection{
		db: db,
	}
}
