package routers

import (
	"e2ee/controllers"
	"e2ee/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetPeerRoutes(router *mux.Router) *mux.Router {
	router.Handle("/peer/{username}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.PeerGet),
		)).Methods("GET")
	router.Handle("/peer",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.PeerNotify),
		)).Methods("POST")

	return router
}
