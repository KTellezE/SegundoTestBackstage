package contexts

import (
	"fmt"
	"log"

	"application/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	UserConfig *config.UserConfig
	DB         *gorm.DB
}

func NewMySQLDB(userConfig *config.UserConfig) (*MySQLDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userConfig.DBUser,
		userConfig.DBPassword,
		userConfig.DBPort,
		userConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Conexión establecida correctamente")

	return &MySQLDB{
		UserConfig: userConfig,
		DB:         db,
	}, nil
}
