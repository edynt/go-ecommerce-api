package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/edynt/go-ecommerce-api/internal/repo"
	"github.com/edynt/go-ecommerce-api/internal/utils/crypto"
	"github.com/edynt/go-ecommerce-api/internal/utils/random"
	"github.com/edynt/go-ecommerce-api/internal/utils/sendto"
	"github.com/edynt/go-ecommerce-api/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUuserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUuserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0 hash Email
	fmt.Println("call this")
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash Email: %s\n", hashEmail)

	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is ::: %d\n", otp)

	// 3. Save OTP in redis with expiration time
	err := us.userAuthRepo.AddOtp(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOtp
	}

	// 4. Send Email OTP
	err = sendto.SendTemplateEmailOtp([]string{email}, "tringuyen20101998bb@gmail.com", "otp-auth.html", map[string]interface{}{
		"Otp": strconv.Itoa(otp),
	})

	if err != nil {
		return response.ErrSendEmailOtp
	}

	return response.ErrCodeSuccess
}
