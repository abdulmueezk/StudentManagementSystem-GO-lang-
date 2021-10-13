package database

import (
	models2 "crud/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Showstudent(w http.ResponseWriter, r *http.Request) {
	var student models2.Student
	db := Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query, usertoken := "SELECT * FROM student WHERE stdemail='"+student.Stdemail+"'", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if usertoken == GenToken {
		resultchk, err := db.Query(querychk)
		if err != nil {
			panic(err)
		}
		defer resultchk.Close()
		var stdcheck models2.Student
		for resultchk.Next() {
			err := resultchk.Scan( &stdcheck.Stdname, &stdcheck.Stdemail, &stdcheck.Stdclass, &stdcheck.Stdage, &stdcheck.Stdcity, &stdcheck.Stdsubject)
			if err != nil {
				panic(err.Error())
			}
		}
		if student.Stdemail == stdcheck.Stdemail {
			result, err := db.Query(query)
			if err != nil {
				panic(err)
			}
			defer result.Close()
			var stud models2.Student
			for result.Next() {
				err := result.Scan(&stud.Stdname, &stud.Stdemail, &stud.Stdclass, &stud.Stdage, &stud.Stdcity, &stud.Stdsubject)
				if err != nil {
					panic(err.Error())
				}
			}
			fmt.Println(stud)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusFound)                    //status code
			json.NewEncoder(w).Encode(stud)                    // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(Validmessage)
		} else {
			var insmess = "No Record found on this Email"
			json.NewEncoder(w).Encode(insmess)
		}
	} else {
		json.NewEncoder(w).Encode(InvalidMessage)
	}
}
