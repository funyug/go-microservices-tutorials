package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *Auth) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
