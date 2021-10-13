package models

type Student struct {
	Token      string `json:"token"`
	Stdid      string `json:"stdid"`
	Stdname    string `json:"stdname"`
	Stdemail   string `json:"stdemail"`
	Stdclass   string `json:"stdclass"`
	Stdage     string `json:"stdage"`
	Stdcity    string `json:"stdcity"`
	Stdsubject string `json:"stdsubject"`
}
