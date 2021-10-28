package micagen

import (
	"github.com/ekokurniadi/micagen/core"
	"gorm.io/gorm"
)

type Micagen struct {
}

func (micagen *Micagen) Generate(db *gorm.DB, model interface{}) {
	core.GenerateAll(db, model)
}
