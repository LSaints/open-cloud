package user

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Init(isRegister bool) error {
	if err := user.validate(isRegister); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(isRegister bool) error {
	if user.Name == "" {
		return errors.New("name field cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email field cannot be empty")
	}
	if isRegister && user.Password == "" {
		return errors.New("password field cannot be empty")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}
