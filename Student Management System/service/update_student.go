package service

import (
	"crud/database"
	"net/http"
)

func UpdateStudnetApiCall()  {
	http.HandleFunc("/updatestudent", database.Updatestudent)
}