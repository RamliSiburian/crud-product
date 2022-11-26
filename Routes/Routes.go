package Routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoute(r)
	ProductRoute(r)
}
