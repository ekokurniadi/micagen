package micagen

import (
	"github.com/ekokurniadi/micagen/core"
	"gorm.io/gorm"
)

func Generate(db *gorm.DB, model interface{}) {
	core.GenerateAll(db, model)
}
