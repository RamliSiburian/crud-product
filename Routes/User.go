package Routes

import (
	"nutecth/Handlers"
	midleware "nutecth/Pkg/Midleware"
	"nutecth/Pkg/Mysql"
	"nutecth/Repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userRepository := Repositories.RepositoryUser(Mysql.DB)
	h := Handlers.HandlerUser(userRepository)

	r.HandleFunc("/User/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/Login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", midleware.Auth(h.CheckAuth)).Methods("GET")
}
