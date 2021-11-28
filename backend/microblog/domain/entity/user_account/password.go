package useraccount

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"microblog/microblog/domain"

	"golang.org/x/crypto/bcrypt"
)

const PASSWORD_VALID_REGEX = "^[[:alpha:]]+$" // ASCII characters
const PASSWORD_MIN_LENGTH = 8
const PASSWORD_MAX_LENGTH = 128

type Password struct {
	hashedPassword string
}

func NewPassword(password string) (*Password, error) {
	if match, _ := regexp.MatchString(PASSWORD_VALID_REGEX, password); !match {
		return nil, fmt.Errorf("%w: The password contains invalid characters", domain.ErrInvalidRequest)
	}

	if len(password) < PASSWORD_MIN_LENGTH {
		return nil, fmt.Errorf("%w: The password must be at least %d characters", domain.ErrInvalidRequest, PASSWORD_MIN_LENGTH)
	}

	if len(password) > PASSWORD_MAX_LENGTH {
		return nil, fmt.Errorf("%w: The password must be less than %d characters", domain.ErrInvalidRequest, PASSWORD_MIN_LENGTH)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return &Password{string(bytes)}, err
}

func (p *Password) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hashedPassword), []byte(password))
	return err == nil
}

func (p *Password) Value() (driver.Value, error) {
	return p.hashedPassword, nil
}

func (p *Password) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		p.hashedPassword = string(v)
		return nil
	}

	return errors.New("failed to scan password")
}
