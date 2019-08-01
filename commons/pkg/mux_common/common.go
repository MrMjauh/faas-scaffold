package mux_common

import (
	"encoding/json"
	"faas-scaffold/commons/pkg/rest"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"runtime/debug"
)

func CreateRouteConfiguration(router * mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(NotFound)
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

func NotFound(w http.ResponseWriter, r *http.Request) {
	// Order is important https://github.com/dimfeld/httptreemux/issues/47
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(rest.HTTP_STATUS_CODE_ERROR)

	resp := rest.Response{
		Data: rest.Error{
			Msg: r.URL.Path + " is not a valid endpoint",
			Code: rest.ERROR_CODE_NOT_FOUND,
		},
	}
	jsonBytes, err := json.Marshal(resp)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonBytes)
}

func WrappedHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			uuid := uuid.Must(uuid.NewV4()).String()
			log.Println("uuid = ", uuid)
			log.Println("Panic recovery with message: ", r)
			log.Println("Stacktrace output")
			log.Println(string(debug.Stack()))

			// Order is important https://github.com/dimfeld/httptreemux/issues/47
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(rest.HTTP_STATUS_CODE_ERROR)

			resp := rest.Response{
				Data: rest.Error{
					Msg: "Oh noe, something went horribly wrong, please show the following UUID = " + uuid + " to the monkeys working on this",
					Code: rest.ERROR_CODE_INTERNAL_ERROR,
				},
			}

			jsonBytes, err := json.Marshal(resp)


			if err != nil {
				log.Fatal(err)
			}

			w.Write(jsonBytes)
		}()
		h.ServeHTTP(w, r)
	})
}
