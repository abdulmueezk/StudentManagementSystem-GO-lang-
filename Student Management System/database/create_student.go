package database

import (
	models2 "crud/models"
	"encoding/json"
	"fmt"
	"net/http"
)
var GenToken string = "123456"
var InvalidMessage = "Invalid Token/Enter Token"
var Validmessage = "Done with your task"
var Emailcheck = "Email already Registerd"
func Createstudent(w http.ResponseWriter, r *http.Request) {

	var student models2.Student
	db := Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}

	query, usertoken := "insert into student (stdname,stdemail,stdclass,stdage,stdcity,stdsubject) values ( '"+student.Stdname+"' ,'"+student.Stdemail+"' ,'"+student.Stdclass+"','"+student.Stdage+"','"+student.Stdcity+"','"+student.Stdsubject+"')", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if usertoken == GenToken {
		resultchk, err := db.Query(querychk)
		if err != nil {
			panic(err)
		}
		defer resultchk.Close()
		var stdcheck models2.Student
		for resultchk.Next() {
			err := resultchk.Scan(&stdcheck.Stdid, &stdcheck.Stdname, &stdcheck.Stdemail, &stdcheck.Stdclass, &stdcheck.Stdage, &stdcheck.Stdcity, &stdcheck.Stdsubject)
			if err != nil {
				panic(err.Error())
			}
		}
		if student.Stdemail != stdcheck.Stdemail {
			result, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusCreated)                  //status code
			json.NewEncoder(w).Encode(student)                 // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(Validmessage)
		} else {
			var insmess = "Email already Registerd"
			json.NewEncoder(w).Encode(insmess)

		}
	} else {
		json.NewEncoder(w).Encode(InvalidMessage)

	}
}
