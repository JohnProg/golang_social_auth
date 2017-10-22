package server

import (
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/golang_social_auth/database"
	"github.com/golang_social_auth/controllers"
	"log"
)

type Server struct {
	Router http.Handler
}

func (s *Server) Initialize() {
	if err := database.ConnectMysql(); err != nil {
		log.Fatal(err.Error())
	}

	var (
		users = controllers.NewUsersController()
	)

	api := rest.NewApi()
	api.Use(rest.DefaultCommonStack...)
	router, err := rest.MakeRouter(
		rest.Post("/signup", users.SignUp),
		rest.Post("/login", users.Login),
		rest.Post("/password_reset", users.PasswordResetByEmail),
		rest.Post("/:token/password_reset", users.PasswordReset),
		rest.Put("/user/me", users.UpdateMe),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	s.Router = api.MakeHandler()
}
