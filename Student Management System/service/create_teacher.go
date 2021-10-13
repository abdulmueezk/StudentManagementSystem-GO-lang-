package service

import (
	"crud/database"
	"net/http"
)

func CreateteacherApiCall()  {
	http.HandleFunc("/createteacher", database.Createteacher)
}