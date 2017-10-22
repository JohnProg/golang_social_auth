package controllers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/golang_social_auth/models"
	"strings"
	"net/http"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseController struct {
	User *models.User
}

func (c *BaseController) Authenticate(w rest.ResponseWriter, r *rest.Request) bool {
	token := r.Request.Header.Get("Authorization")
	if token == "" {
		token = r.Request.URL.Query().Get("access_token")
	} else {
		token = strings.TrimPrefix(token, "token ")
	}

	u, err := UserHandler.ByToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.WriteJson(Error{
			Code:    "err:unknown",
			Message: "Login required",
		})
		return false
	}
	c.User = u
	return true
}
