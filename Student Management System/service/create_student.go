package service

import (
	"crud/database"
	"net/http"
)

func CreateStudnetApiCall()  {
	http.HandleFunc("/createstudent", database.Createstudent)
}
