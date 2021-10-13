package handler

import (
	"crud/service"
	"net/http"
)

func LoginteacherApiCall(w http.ResponseWriter, r *http.Request)  {
 	service.LoginteacherApiCall()
}