package handler

import (
 "crud/service"
 "net/http"
)

func CreateteacherApiCall(w http.ResponseWriter, r *http.Request)  {
 service.CreateteacherApiCall()
}