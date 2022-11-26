package Routes

import (
	"nutecth/Handlers"
	midleware "nutecth/Pkg/Midleware"
	"nutecth/Pkg/Mysql"
	"nutecth/Repositories"

	"github.com/gorilla/mux"
)

func ProductRoute(r *mux.Router) {
	productRepository := Repositories.RepositoryProduct(Mysql.DB)
	h := Handlers.HandlerProduct(productRepository)

	r.HandleFunc("/Products", h.FindProduct).Methods("GET")
	r.HandleFunc("/Product/{id}", h.GetProductById).Methods("GET")
	r.HandleFunc("/ProductByUser/{user_id}", h.GetProductByUser).Methods("GET")
	r.HandleFunc("/Product", midleware.Auth(midleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/Product/{id}", midleware.UploadFile(h.UpdateProduct)).Methods("PATCH")
	r.HandleFunc("/Product/{id}", midleware.Auth(h.DeleteProduct)).Methods("DELETE")
}
