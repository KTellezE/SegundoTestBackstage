package models

import (
	"application/config"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	LastName string `gorm:"size:255"`
}

func (User) TableName() string {
	tableName, err := config.GetEnvVariable("DB_TABLE")
	if err != nil {
		log.Fatalf("No se pudo obtener el valor del nombre de la tabla: %v", err)
	}
	return tableName
}
