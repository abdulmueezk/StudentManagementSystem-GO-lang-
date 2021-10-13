package handler

import (
	"crud/service"
	"net/http"
)

func ReadStudnetApiCall(w http.ResponseWriter, r *http.Request)  {
	service.ReadStudnetApiCall()
}