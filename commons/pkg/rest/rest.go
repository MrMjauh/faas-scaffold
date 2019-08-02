package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/satori/go.uuid"
)

type Response struct {
	Data interface{}
}

type Error struct {
	Msg string
	Code uint32
}

const (
	HTTP_STATUS_CODE_OK = 200
	HTTP_STATUS_CODE_ERROR = 400
)

const (
	ERROR_CODE_VALIDATION_ERROR = 7
	ERROR_CODE_NOT_FOUND = 400
	ERROR_CODE_INTERNAL_ERROR = 401
)

func GeneralErrorResponse(code uint32, msg string) Response {
	return Response{
		Data: Error{
			Msg:  msg,
			Code: code,
		},
	}
}

func InternalErrorResponse() (Response, string) {
	uuid := GetUUID()

	return Response{
		Data: Error{
			Msg:  "Something went horribly wrong, check UUID = " + uuid + " for more information",
			Code: ERROR_CODE_INTERNAL_ERROR,
		},
	}, uuid
}

func ValidationErrorResponse(fieldName string) Response {
	return GeneralErrorResponse(ERROR_CODE_VALIDATION_ERROR, "Field " + fieldName + " is either not valid or not present")
}

func WriteJsonError(w http.ResponseWriter, response Response) {
	jsonBytes, err := json.Marshal(response)

	if err != nil {
		log.Println("Could not marshal error response, this should never happen")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(HTTP_STATUS_CODE_ERROR)
	w.Write(jsonBytes)
}

func WriteJsonResponse(w http.ResponseWriter, data interface{}) {
	jsonBytes, err := json.Marshal(data)

	if err != nil {
		resp, uuid := InternalErrorResponse()
		log.Println(uuid)
		log.Println(err)
		WriteJsonError(w, resp)
		return
	}

	// Order is important https://github.com/dimfeld/httptreemux/issues/47
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(HTTP_STATUS_CODE_OK)
	w.Write(jsonBytes)
}

func GetUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}