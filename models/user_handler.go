package models

import (
	"github.com/golang_social_auth/settings"
	"github.com/golang_social_auth/database"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"errors"
	"github.com/satori/go.uuid"
	"log"
	"github.com/golang_social_auth/mailers"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

type User struct {
	ID        int
	Type      string
	Value     string
	Password  string
	Token     string
	FirstName string
	LastName  string
	Username  string
}

const (
	ProviderTypeEmail   = "email"
	ProvideTypeLinkedIn = "linkedin"
)

// SignUp registers a new user
// If user is exist will return error
// If type email, will be generated bcrypt from password for security reason
func (h *userHandler) SignUp(u *User) (*User, error) {

	u.Type = strings.ToLower(u.Type)
	u.Value = strings.ToLower(u.Value)
	u.Token = uuid.NewV4().String()

	if u.Type != "" {
		if !h.TypeIsValid(u.Type) {
			return nil, errors.New("type is invalid")
		}
	}

	if h.UserIsAvailable(u.Type, u.Value) {
		return nil, errors.New("user is in use")
	}

	var (
		hpass []byte
		err   error
	)
	if u.Type == ProviderTypeEmail {
		hpass, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
	}
	u.Password = string(hpass)

	if err := database.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	u.Password = ""
	return u, nil
}

func (h *userHandler) Login(u *User) (*User, error) {
	u.Type = strings.ToLower(u.Type)
	u.Value = strings.ToLower(u.Value)
	u.Token = uuid.NewV4().String()

	if u.Type != "" {
		if !h.TypeIsValid(u.Type) {
			return nil, errors.New("type is invalid")
		}
	}

	user, err := h.ByTypeAndValue(u.Type, u.Value)
	if err != nil {
		return nil, err
	}

	if u.Type == ProviderTypeEmail {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
			return nil, errors.New("password invalid")
		}
	}

	user.Password = ""
	return user, nil
}

func (h *userHandler) UpdateMe(u *User) error {
	return database.DB.Model(&u).UpdateColumns(User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
	}).Error
}

func (h *userHandler) PasswordResetByEmail(email, host string) error {
	u, err := h.ByTypeAndValue(ProviderTypeEmail, email)
	if err != nil {
		return err
	}
	go func() {
		h.sendVerification(host, u)
	}()
	return nil
}

func (h *userHandler) PasswordReset(password, token string) error {
	u, err := h.ByToken(token)
	if err != nil {
		return errors.New("token is not valid")
	}

	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Token = uuid.NewV4().String()
	u.Password = string(hpass)

	return database.DB.Model(&u).UpdateColumns(&User{
		Token:    u.Token,
		Password: u.Password,
	}).Error
}

func (h *userHandler) TypeIsValid(typ string) bool {
	return ProviderTypeEmail == typ || ProvideTypeLinkedIn == typ
}

func (h *userHandler) UserIsAvailable(typ, value string) bool {
	count := 0
	database.DB.Model(&User{}).Where("LOWER(type) = ? AND LOWER(value) = ?", strings.ToLower(typ), strings.ToLower(value)).Limit(1).Count(&count)
	return count != 0
}

func (h *userHandler) ByTypeAndValue(typ, value string) (*User, error) {
	u := &User{}
	if err := database.DB.Where("LOWER(type) = ? AND LOWER(value) = ?", strings.ToLower(typ), strings.ToLower(value)).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (h *userHandler) ByToken(token string) (*User, error) {
	u := &User{}
	if err := database.DB.Where("token = ?", token).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (h *userHandler) sendVerification(host string, u *User) error {
	binding := struct {
		User     *User
		LiveHost string
	}{
		User:     u,
		LiveHost: host,
	}

	body, err := mailers.RenderTemplate("en_us/verification", binding)
	if err != nil {
		return err
	}

	if err := mailers.Send(
		u.Value,
		settings.Config.I18n.Get("en_us", "verification.email_subject"),
		body,
	); err != nil {
		log.Printf("Email fail to %s with error:%v\n", u.Value, err.Error())
		return err
	}

	log.Printf("Email send to %s with token %s\n", u.Value, u.Token)
	return nil
}
