package repo

import (
	"github.com/edynt/go-ecommerce-api/global"
	"github.com/edynt/go-ecommerce-api/internal/model"
)

type IUuserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUuserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

	return row != NumberNull
}

func NewUserRepository() IUuserRepository {
	return &userRepository{}
}
