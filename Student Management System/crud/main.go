package main

import (
	"crud/handler"
	"fmt"
	"net/http"

)

func main() {
	http.HandleFunc("/createteacher", handler.CreateteacherApiCall)
	http.HandleFunc("/loginteacher", handler.LoginteacherApiCall)
	http.HandleFunc("/createstudent", handler.CreateStudnetApiCall)
	http.HandleFunc("/updatestudent", handler.UpdateStudnetApiCall)
	http.HandleFunc("/deletestudent", handler.DeleteStudnetApiCall)
	http.HandleFunc("/readstudent", handler.ReadStudnetApiCall)
	fmt.Println(http.ListenAndServe(":8081", nil))
}
