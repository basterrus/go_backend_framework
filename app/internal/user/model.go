package user

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             int       `json:"id"`
	Uuid           uuid.UUID `json:"uuid"`
	Username       string    `json:"username"`
	FirstName      string    `json:"first_name,omitempty"`
	LastName       string    `json:"last_name,omitempty"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	RepeatPassword string    `json:"-"`
}

type CreateUserDTO struct {
	Id             int           `json:"-"`
	Uuid           uuid.UUID     `json:"uuid"`
	Username       string        `json:"username"`
	FirstName      string        `json:"first_name,omitempty"`
	LastName       string        `json:"last_name,omitempty"`
	Email          string        `json:"email"`
	Password       string        `json:"password"`
	RepeatPassword string        `json:"repeat_password"`
	CreatedAt      time.Duration `json:"created_at"`
}

type UpdateUserDTO struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
}

func NewUser(dto CreateUserDTO) User {
	return User{
		Username:       dto.Username,
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Password:       dto.Password,
		RepeatPassword: dto.RepeatPassword,
	}
}

func UpdatedUser(dto UpdateUserDTO) User {
	return User{
		Username:  dto.Username,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("password does not match")
	}
	return nil
}

func (u *User) GeneratePasswordHash() error {
	pwd, err := generatePasswordHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwd
	return nil
}
