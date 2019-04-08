package routerHandlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type Profiletpl struct {
	Cpyname        string
	Cpycategory    string
	Cpydescription string
	Username       string
	IsOnline       bool
}

func CpyProfile(w http.ResponseWriter, r *http.Request) {
	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}

	var databaseCpyName sql.NullString
	var databaseCpyCategory sql.NullString
	var databaseDescription sql.NullString

	row := DB.QueryRow("select companyname,category,discription from cpyusers where username = ?", session.Username)
	err := row.Scan(&databaseCpyName, &databaseCpyCategory, &databaseDescription)
	if err != nil {
		panic(err.Error())
		fmt.Println(err)
		return
	}

	tplhandler := Profiletpl{databaseCpyName.String, databaseCpyCategory.String, databaseDescription.String, session.Username, session.Connection}

	profiletpl, err := template.ParseFiles(PPath+"/views/cpyProfile.html", PPath+"/views/hnavbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		panic(err.Error())
		fmt.Println(err)
		return
	}

	err = profiletpl.Execute(w, tplhandler)
	if err != nil {
		panic(err.Error())
		fmt.Println(err)
		return
	}
}
