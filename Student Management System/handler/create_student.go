package handler

import (
	"crud/service"
	"net/http"
)

func CreateStudnetApiCall(w http.ResponseWriter, r *http.Request)  {
	service.CreateStudnetApiCall()
}
