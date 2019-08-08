package handler

import (
	"github.com/MrMjauh/faas-scaffold/commons/pkg/rest"
	"github.com/MrMjauh/faas-scaffold/multiply/internal/pkg/config"
	"github.com/MrMjauh/faas-scaffold/multiply/internal/pkg/service"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type AnswerReturn struct {
	Result int64
	ServiceID string
}

func MultiplyHandler(w http.ResponseWriter, r * http.Request) {
	xStr := r.URL.Query().Get("x")
	yStr := r.URL.Query().Get("y")
	x, err := strconv.ParseInt(xStr, 10, 64)
	if err != nil {
		rest_common.WriteJsonError(w, rest_common.ValidationErrorResponse("x"))
		return
	}
	y, err := strconv.ParseInt(yStr, 10, 64)
	if err != nil {
		rest_common.WriteJsonError(w, rest_common.ValidationErrorResponse("y"))
		return
	}

	res, ok := service.Multiply(x, y)
	if !ok {
		rest_common.WriteJsonError(w, rest_common.GeneralErrorResponse(config.ERROR_CODE_COMPUTE_OVERFLOW, "Overflow"))
		return
	}

	rest_common.WriteJsonResponse(w, AnswerReturn{
			Result: res,
			ServiceID: serviceId(),
	})
}

func AdditionHandler(w http.ResponseWriter, r * http.Request) {
	xStr := r.URL.Query().Get("x")
	yStr := r.URL.Query().Get("y")
	x, err := strconv.ParseInt(xStr, 10, 64)
	if err != nil {
		rest_common.WriteJsonError(w, rest_common.ValidationErrorResponse("x"))
		return
	}
	y, err := strconv.ParseInt(yStr, 10, 64)
	if err != nil {
		rest_common.WriteJsonError(w, rest_common.ValidationErrorResponse("y"))
		return
	}

	res, ok := service.Add(x, y)
	if !ok {
		rest_common.WriteJsonError(w, rest_common.GeneralErrorResponse(config.ERROR_CODE_COMPUTE_OVERFLOW, "Overflow"))
		return
	}

	rest_common.WriteJsonResponse(w, AnswerReturn{
			Result: res,
			ServiceID: serviceId(),
	})
}

func PanicHandle(w http.ResponseWriter, r * http.Request) {
	panic("Paaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaanic")
}

func serviceId() string {
	out, _ := exec.Command("cat", "/etc/hostname").Output()
	return strings.TrimSuffix(strings.TrimSpace(string(out)), "\n")
}