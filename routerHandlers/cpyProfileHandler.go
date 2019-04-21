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
	requestHeader := FormatRequest(r)
	fmt.Print("\n")
	fmt.Println(requestHeader)
	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		category := r.FormValue("category")
		description := r.FormValue("description")
		fmt.Println("description is: ", description)

		updateProfilehandler, err := DB.Prepare("UPDATE cpyusers SET category = ?, description = ? WHERE username = ?")
		if err != nil {
			panic(err.Error())
		}
		_, err = updateProfilehandler.Exec(category, description, session.Username);
		http.Redirect(w, r, "/cpyIndex/profile", http.StatusSeeOther)
	}

	var databaseCpyName sql.NullString
	var databaseCpyCategory sql.NullString
	var databaseDescription sql.NullString

	row := DB.QueryRow("select companyname,category,description from cpyusers where username = ?", session.Username)
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
