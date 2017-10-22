package controllers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/golang_social_auth/models"
	"net/http"
)

var UserHandler = models.NewUserHandler()

type UsersController struct {
	BaseController
}

type User struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Token     string `json:"token"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (c *UsersController) SignUp(w rest.ResponseWriter, r *rest.Request) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{}

	err := r.DecodeJsonPayload(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: "An unknown error occurred",
		})
		return
	}

	u := &models.User{
		Type:     in.Type,
		Value:    in.Value,
		Password: in.Password,
	}

	user, err := UserHandler.SignUp(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: err.Error(),
		})
		return
	}

	w.WriteJson(map[string]User{"user": makeUser(user)})
}

func (c *UsersController) Login(w rest.ResponseWriter, r *rest.Request) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{}

	err := r.DecodeJsonPayload(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: "An unknown error occurred",
		})
		return
	}

	u := &models.User{
		Type:     in.Type,
		Value:    in.Value,
		Password: in.Password,
	}

	user, err := UserHandler.Login(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: err.Error(),
		})
		return
	}

	w.WriteJson(map[string]User{"user": makeUser(user)})
}

func (c *UsersController) UpdateMe(w rest.ResponseWriter, r *rest.Request) {
	if !c.Authenticate(w, r) {
		return
	}

	in := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	}{}

	err := r.DecodeJsonPayload(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: "An unknown error occurred",
		})
		return
	}

	c.User.FirstName = in.FirstName
	c.User.LastName = in.LastName
	c.User.Username = in.Username

	if err := UserHandler.UpdateMe(c.User); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: err.Error(),
		})
		return
	}

	w.WriteJson(map[string]User{"user": makeUser(c.User)})
}

func (c *UsersController) PasswordResetByEmail(w rest.ResponseWriter, r *rest.Request) {
	in := struct {
		Email string `json:"email"`
	}{}

	err := r.DecodeJsonPayload(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: "An unknown error occurred",
		})
		return
	}

	if err := UserHandler.PasswordResetByEmail(in.Email, r.Host); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *UsersController) PasswordReset(w rest.ResponseWriter, r *rest.Request) {
	in := struct {
		Password string `json:"password"`
	}{}

	err := r.DecodeJsonPayload(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: "An unknown error occurred",
		})
		return
	}
	if err := UserHandler.PasswordReset(in.Password, r.PathParam("token")); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func makeUser(u *models.User) User {
	return User{
		Type:      u.Type,
		Value:     u.Value,
		Token:     u.Token,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
	}
}
