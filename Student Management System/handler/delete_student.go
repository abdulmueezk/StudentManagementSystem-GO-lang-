package handler

import (
	"crud/service"
	"net/http"
)

func DeleteStudnetApiCall(w http.ResponseWriter, r *http.Request)  {
	service.DeleteStudnetApiCall()
}