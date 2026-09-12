package services

import (
	"money_management/dtos"
	"money_management/models"
	"money_management/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(data dtos.LoginRequest) (string, int64, error)
	Register(data dtos.RegisterRequest) (string, error)
	Logout(token string) error
	generateToken(user models.User) (string, int64, error)
}

type authService struct {
	userService UserService
}

func (a authService) generateToken(user models.User) (string, int64, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
		exp int64
	)
	exp = time.Now().Add(time.Hour * 24 * 7).Unix()
	key = []byte(utils.GetJWTSecretKey())
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.Id,
			"iss":     user.Username,
			"sub":     user.FullName,
			"iat":     time.Now().Unix(),
			"exp":     exp,
		})
	s, err := t.SignedString(key)
	if err != nil {
		return "", exp, err
	}
	return s, exp, nil
}

func (a authService) Register(data dtos.RegisterRequest) (string, error) {
	user, _ := a.userService.FindByUsername(data.Username)
	if user.Id != "" {
		return "", utils.ErrUserAlreadyExist
	}
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	userCreate := models.User{
		Username:   data.Username,
		FullName:   data.FullName,
		Password:   string(hashedPasswordBytes),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		VerifiedAt: nil,
	}
	result, err := a.userService.RegisterUser(userCreate)
	if err != nil {
		return "", err
	}
	return result.Username, nil
}

func (a authService) Login(data dtos.LoginRequest) (string, int64, error) {
	user, err := a.userService.FindByUsername(data.Username)
	if err != nil {
		return "", 0, utils.ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", 0, err
	}
	token, eat, err := a.generateToken(user)
	if err != nil {
		return "", eat, err
	}
	return token, eat, nil
}

func (a authService) Logout(token string) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(userService UserService) AuthService {
	return authService{
		userService: userService,
	}
}
