package real_emails_srv

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ReplaceRealEmail struct {
	fx.In
	Db *gorm.DB
}

func (rre ReplaceRealEmail) Replace(
	user models.UserModel,
	oldEmail string,
	newEmail string,
) {
	model := models.RealAddress{}
	uQuery := "user_id = ?"
	realAddrQuery := "real_address = ?"

	rre.Db.Model(&model).Where(uQuery, user.Id).Where(realAddrQuery, oldEmail).Update(realAddrQuery, newEmail)
}
