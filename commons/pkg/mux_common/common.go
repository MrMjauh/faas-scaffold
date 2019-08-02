package mux_common

import (
	"encoding/json"
	"faas-scaffold/commons/pkg/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime/debug"
)

func CreateRouteConfiguration(router * mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	router.Use(RestMiddleware)
}

func CreateRouteForVersion(router * mux.Router, version string) *mux.Router {
	return router.PathPrefix("/api/" + version).Subrouter()
}

func RestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Order is important https://github.com/dimfeld/httptreemux/issues/47
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(rest.HTTP_STATUS_CODE_ERROR)

	rest.WriteJsonError(w, rest.Response{
			Data: rest.Error{
				Msg:  r.URL.Path + " is not a valid endpoint",
				Code: rest.ERROR_CODE_NOT_FOUND,
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

			resp, uuid := rest.InternalErrorResponse()
			log.Println("uuid = ", uuid)
			log.Println("Panic recovery with message: ", r)
			log.Println("Stacktrace output")
			log.Println(string(debug.Stack()))

			rest.WriteJsonResponse(w, resp)
		}()
		h.ServeHTTP(w, r)
	})
}
