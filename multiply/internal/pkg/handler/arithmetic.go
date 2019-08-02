package handler

import (
	"faas-scaffold/commons/pkg/rest"
	"faas-scaffold/multiply/internal/pkg/config"
	"faas-scaffold/multiply/internal/pkg/service"
	"net/http"
	"strconv"
)

type AnswerReturn struct {
	Result int64
}

func MultiplyHandler(w http.ResponseWriter, r * http.Request) {
	xStr := r.URL.Query().Get("x")
	yStr := r.URL.Query().Get("y")
	x, err := strconv.ParseInt(xStr, 10, 64)
	if err != nil {
		rest.WriteJsonError(w, rest.ValidationErrorResponse("x"))
		return
	}
	y, err := strconv.ParseInt(yStr, 10, 64)
	if err != nil {
		rest.WriteJsonError(w, rest.ValidationErrorResponse("y"))
		return
	}

	res, ok := service.Multiply(x, y)
	if !ok {
		rest.WriteJsonError(w, rest.GeneralErrorResponse(config.ERROR_CODE_COMPUTE_OVERFLOW, "Overflow"))
		return
	}

	rest.WriteJsonResponse(w, rest.Response{
		Data: AnswerReturn{
			Result: res,
		},
	})
}

func AdditionHandler(w http.ResponseWriter, r * http.Request) {
	xStr := r.URL.Query().Get("x")
	yStr := r.URL.Query().Get("y")
	x, err := strconv.ParseInt(xStr, 10, 64)
	if err != nil {
		rest.WriteJsonError(w, rest.ValidationErrorResponse("x"))
		return
	}
	y, err := strconv.ParseInt(yStr, 10, 64)
	if err != nil {
		rest.WriteJsonError(w, rest.ValidationErrorResponse("y"))
		return
	}

	res, ok := service.Add(x, y)
	if !ok {
		rest.WriteJsonError(w, rest.GeneralErrorResponse(config.ERROR_CODE_COMPUTE_OVERFLOW, "Overflow"))
		return
	}

	rest.WriteJsonResponse(w, rest.Response{
		Data: AnswerReturn{
			Result: res,
		},
	})
}