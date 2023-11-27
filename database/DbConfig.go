package database

import (
	"github.com/ekart/payments/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var paymentDb *gorm.DB

func Connection(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Databse not connected..")
	}
	paymentDb = db

	paymentDb.AutoMigrate(model.Payment{})
}

func PayemntDB() *gorm.DB {
	return paymentDb
}
