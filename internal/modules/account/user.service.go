package account

import (
	"go-api/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET = []byte("your_secret")

type UserService struct {
	UserRepo *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{UserRepo: r}
}

// REGISTER
func (s *UserService) Register(email string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	userExist, err := s.UserRepo.FindByEmail(email)
	if userExist != nil {
		return utils.NewApiError(400, "User already exist")
	}

	return s.UserRepo.Create(email, string(hash))
}

// LOGIN
func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", utils.NewApiError(401, "invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", utils.NewApiError(401, "invalid credentials")
	}

	// JWT payload
	claims := jwt.MapClaims{
		"user_uuid": user.UserUUID,
		"role_name": user.RoleName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}

	return signed, nil
}
