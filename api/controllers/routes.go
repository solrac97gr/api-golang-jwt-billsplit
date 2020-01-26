package controllers

import "github.com/solrac97gr/api-golang-jwt/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	s.Router.HandleFunc("/usersemail/{email}", middlewares.SetMiddlewareJSON(s.GetUserEmail)).Methods("GET")

	//Pay routes
	s.Router.HandleFunc("/pays", middlewares.SetMiddlewareJSON(s.CreatePay)).Methods("POST")
	s.Router.HandleFunc("/pays", middlewares.SetMiddlewareJSON(s.GetPays)).Methods("GET")
	s.Router.HandleFunc("/pays/{id}", middlewares.SetMiddlewareJSON(s.GetPay)).Methods("GET")
	s.Router.HandleFunc("/paysbyuser/{id}", middlewares.SetMiddlewareJSON(s.GetPaysByUserId)).Methods("GET")
	s.Router.HandleFunc("/pays/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePay))).Methods("PUT")
	s.Router.HandleFunc("/pays/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePay)).Methods("DELETE")
}
