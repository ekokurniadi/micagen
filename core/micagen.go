package core

import "gorm.io/gorm"

func GenerateAll(db *gorm.DB, model interface{}) {
	GenerateTable(db, model)
	CreateStructInput(model)
	CreateRepository(db, model)
	CreateService(db, model)
	CreateFormatter(db, model)
}
