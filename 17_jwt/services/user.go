package services

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserService struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserService{db: db}
}

func (u *UserService) New(user User) (string, error) {
	if user.Password == "" {
		return "", errors.New("password empty")
	}

	user.Id = uuid.NewString()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return "", errors.New(err.Error())
	}
	user.Password = string(hash)

	u.db.Create(&user)
	return user.Id, nil
}

func (u *UserService) Authentication(login string, password string) (*struct {
	AccessToken  string
	RefreshToken string
}, error) {
	user := new(User)
	u.db.Where("login = ?", login).Find(user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	accessToken, err := NewAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := NewRefreshToken()
	if err != nil {
		return nil, err
	}

	tokens := struct {
		AccessToken  string
		RefreshToken string
	}{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &tokens, nil
}
