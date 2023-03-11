package DB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/Model"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(ConfigureDB())
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Model.Todo{})
	fmt.Println("Successfully migrated")

	return db
}

func ConfigureDB() (gorm.Dialector, gorm.Option) {
	HOST := "localhost"
	DBNAME := "gorm"
	PORT := "5432"
	SSLMODE := "disable"
	TIMEZONE := "Asia/Tokyo"
	dsn :=
		"host=" + HOST +
			" dbname=" + DBNAME +
			" port=" + PORT +
			" sslmode=" + SSLMODE +
			" TimeZone=" + TIMEZONE

	return postgres.Open(dsn), &gorm.Config{}
}
