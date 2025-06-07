package repo

import (
	"fmt"
	"time"

	"github.com/edynt/go-ecommerce-api/global"
)

type IUserAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOtp implements IUserAuthRepository.
func (u *userAuthRepository) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
