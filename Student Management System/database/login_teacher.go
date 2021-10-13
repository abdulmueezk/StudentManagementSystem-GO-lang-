package database

import (
	models2 "crud/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Teacherlogin(w http.ResponseWriter, r *http.Request) {
	var teacher models2.Teacher
	gentrateToken := "Token 123456"
	logresult := "Successfull Login"
	db := Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query := "SELECT * FROM teacher WHERE tecemail= '" + teacher.Tecemail + "' AND tecpassword= '" + teacher.Tecpassword + "'"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	var teccheck models2.Teacher
	for result.Next() {
		err := result.Scan( &teccheck.Tecname, &teccheck.Tecemail, &teccheck.Tecpassword)
		if err != nil {
			panic(err.Error())
		}
	}
	if teacher.Tecemail == teccheck.Tecemail && teacher.Tecpassword == teccheck.Tecpassword {
		json.NewEncoder(w).Encode(gentrateToken)
		json.NewEncoder(w).Encode(logresult)
		json.NewEncoder(w).Encode(teccheck) // print values in post man to cheack what vvalue store
	} else {
		var instantMessage = "Invalid Email or Password"
		json.NewEncoder(w).Encode(instantMessage)
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json") //show meaning full massage
	w.WriteHeader(http.StatusFound)                    //status code
}
