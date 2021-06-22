package admin

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/theplant/go-que-admin/models"
)

func ConnectDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DB_PARAMS"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	err = db.AutoMigrate(
		// &models.GoqueJob{},
		&models.TestUser{},
	).Error
	if err != nil {
		panic(err)
	}
	return
}
