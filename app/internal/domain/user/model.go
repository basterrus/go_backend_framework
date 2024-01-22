package user

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id           uint   `json:"-" db:"id"`
	RoleId       int    `json:"role_id" `
	Uuid         string `json:"uuid"`
	Username     string `json:"username" `
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	//AccessToken  string `json:"access_token"`
	//RefreshToken string `json:"refresh_token"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Id             uint          `json:"-"`
	RoleId         int           `json:"role_id"`
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

type SignIn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(dto CreateUserDTO) User {
	return User{
		Username:  dto.Username,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
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

func (u *User) GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}
