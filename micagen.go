package micagen

import "gorm.io/gorm"

type Micagen struct {
}

func (micagen *Micagen) GenerateAll(db *gorm.DB, model interface{}) {
	GenerateTable(db, model)
	CreateStructInput(model)
	CreateRepository(db, model)
	CreateService(db, model)
	CreateFormatter(db, model)
}
