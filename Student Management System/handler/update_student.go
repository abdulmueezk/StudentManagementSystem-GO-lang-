package handler

import (
	"crud/service"
	"net/http"
)

func UpdateStudnetApiCall(w http.ResponseWriter, r *http.Request)  {
	service.UpdateStudnetApiCall()
}