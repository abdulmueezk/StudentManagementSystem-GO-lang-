package service

import (
	"crud/database"
	"net/http"
)

func ReadStudnetApiCall()  {
	http.HandleFunc("/createstudent", database.Showstudent)

}