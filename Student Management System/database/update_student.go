package database

import (
	models2 "crud/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Updatestudent(w http.ResponseWriter, r *http.Request) {
	var student models2.Student
	db := Sqlclient()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query, userToken := "UPDATE student SET stdname='"+student.Stdname+"', stdclass='"+student.Stdclass+"', stdcity='"+student.Stdcity+"', stdage='"+student.Stdage+"', stdsubject='"+student.Stdsubject+"' WHERE stdemail='"+student.Stdemail+"'", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if userToken == GenToken {
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
			result, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusFound)                    //status code
			json.NewEncoder(w).Encode(student)                 // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(Validmessage)
		} else {
			var instantMessage = "Email Not Valid"
			json.NewEncoder(w).Encode(instantMessage)
		}
	} else {
		json.NewEncoder(w).Encode(InvalidMessage)
	}
}
