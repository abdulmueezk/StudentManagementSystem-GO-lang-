package service

import (
	"crud/database"
	"net/http"
)

func LoginteacherApiCall()  {
	http.HandleFunc("/teacherlogin", database.Teacherlogin)
}