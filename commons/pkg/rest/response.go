package rest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data interface{}
}

func GeneralErrorResponse(code uint32, msg string) []byte {
	resp := Response{
		Data: Error{
			Msg:  msg,
			Code: code,
		},
	}

	jsonBytes, _ := json.Marshal(resp)
	return jsonBytes
}

func ValidationErrorResponse(fieldName string) []byte {
	return GeneralErrorResponse(ERROR_CODE_VALIDATION_ERROR, "Field " + fieldName + " is either not valid or not present")
}

func WriteError(w http.ResponseWriter, data []byte) {
	w.WriteHeader(HTTP_STATUS_CODE_ERROR)
	w.Write(data)
}