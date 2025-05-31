package services

import (
	"github.com/edynt/go-ecommerce-api/internal/repo"
	"github.com/edynt/go-ecommerce-api/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUuserRepository
}

func NewUserService(userRepo repo.IUuserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists

	}
	return 0
}
