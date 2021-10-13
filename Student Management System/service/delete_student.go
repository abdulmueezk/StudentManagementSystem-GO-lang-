package service

import (
	"crud/database"
	"net/http"
)

func DeleteStudnetApiCall()  {
	http.HandleFunc("/createstudent", database.Deletestudent)

}