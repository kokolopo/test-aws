package database

import (
	"buat_docker/user"

	"buat_docker/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	config := config.InitConfiguration()

	//dsn := "username:password@tcp(host:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&user.User{})
}
