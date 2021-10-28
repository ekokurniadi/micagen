package micagen

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Micagen struct {
}

func (micagen *Micagen) GenerateAll(db *gorm.DB, model interface{}) {
	GenerateTable(db, model)
	CreateStructInput(model)
	CreateRepository(db, model)
	CreateService(db, model)
	CreateFormatter(db, model)
	CreateHandler(db, model)
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func (micagen *Micagen) ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func (micagen *Micagen) FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors

}
