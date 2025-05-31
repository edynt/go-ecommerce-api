package repo

type IUuserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUuserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUuserRepository {
	return &userRepository{}
}
