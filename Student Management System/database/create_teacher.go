package database

import (
	models2 "crud/models"
	"encoding/json"
	"fmt"
	"net/http"
)



func Createteacher(w http.ResponseWriter, r *http.Request) {
	var teacher models2.Teacher
	db := Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query := "insert into teacher (tecname,tecemail,tecpassword) values ( '" + teacher.Tecname + "' ,'" + teacher.Tecemail + "','" + teacher.Tecpassword + "')"
	querychk := "select * from teacher where tecemail='" + teacher.Tecemail + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		panic(err)
	}
	defer resultchk.Close()
	var teccheck models2.Teacher
	for resultchk.Next() {
		err := resultchk.Scan( &teccheck.Tecname, &teccheck.Tecemail, &teccheck.Tecpassword)
		if err != nil {
			panic(err.Error())
		}
	}
	if teacher.Tecemail != teccheck.Tecemail {
		result, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		defer db.Close()
		w.Header().Set("Content-Type", "application/json") //show meaning full massage
		w.WriteHeader(http.StatusCreated)                  //status code
		json.NewEncoder(w).Encode(teacher)                 // print values in post man to cheack what vvalue store
		json.NewEncoder(w).Encode(Validmessage)
	} else {
		json.NewEncoder(w).Encode(Emailcheck)
	}
}