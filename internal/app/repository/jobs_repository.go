package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type JobsRepository struct {
	fx.In
	Db *gorm.DB
}

func (jbRep JobsRepository) Count() int64 {
	var count int64
	jbRep.Db.Model(models.Jobs{}).Count(&count)
	return count
}
