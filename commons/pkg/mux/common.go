package mux_common

import (
	"github.com/MrMjauh/faas-scaffold/commons/pkg/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime/debug"
)

func CreateRoutingTemplate() *mux.Router {
	router := mux.NewRouter()
	SetupRouterConfiguration(router)
	return router
}

func SetupRouterConfiguration(router * mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	router.Use(RestMiddleware)
}

func CreateAPIRoute(router * mux.Router, version string) *mux.Router {
	return router.PathPrefix("/api/" + version).Subrouter()
}

func RestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	rest_common.WriteJsonError(w, rest_common.Response{
			Data: rest_common.Error{
				Msg:  r.URL.Path + " is not a valid endpoint",
				Code: rest_common.ERROR_CODE_NOT_FOUND,
			},
	})
}

func WrappedHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			error, uuid := rest_common.InternalErrorResponse()
			log.Println("uuid = ", uuid)
			log.Println("Panic recovery with message: ", r)
			log.Println("Stacktrace output")
			log.Println(string(debug.Stack()))

			rest_common.WriteJsonError(w, error)
		}()
		h.ServeHTTP(w, r)
	})
}
