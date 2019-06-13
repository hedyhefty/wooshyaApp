package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"wooshyaApp/Models"
)

type stdProfiletpl struct {
	Student  Models.StdUserModel
	Username string
	IsOnline bool
}

func StdProfile(w http.ResponseWriter, r *http.Request) {
	requestHeader := FormatRequest(r)
	fmt.Print("\n")
	fmt.Println(requestHeader)
	IsOnline, session := CheckLogin(Student, r)
	fmt.Println(IsOnline)
	if !IsOnline {
		//http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		http.Redirect(w, r, "/stdMessage?mtype=14", 301)
		return
	}

	if r.Method == "POST" {
		skills := GetFromValue(r, "skills")
		hobbies := GetFromValue(r, "hobbies")
		updateProfilehandler, err := DB.Prepare("UPDATE stdusers SET skills = ?, hobbies = ? WHERE username = ?")
		if err != nil {
			panic(err.Error())
		}
		_, err = updateProfilehandler.Exec(skills, hobbies, session.Username);
		http.Redirect(w, r, "/stdProfile", http.StatusSeeOther)
	}

	var databaseFirstName sql.NullString
	var databaseLastName sql.NullString
	var databaseCollege sql.NullString
	var databaseSkills sql.NullString
	var databaseHobbies sql.NullString

	row := DB.QueryRow("select firstname,lastname,collegename,skills,hobbies from stdusers where username=?", session.Username)
	err := row.Scan(&databaseFirstName, &databaseLastName, &databaseCollege, &databaseSkills, &databaseHobbies)
	if ErrorHandler(w, err, QueryError, 500) {
		http.Redirect(w, r, "/stdMessage?mtype=100", 301)
	}

	student := Models.StdUserModel{FirstName: databaseFirstName.String, LastName: databaseLastName.String, CollegeName: databaseCollege.String, Skills: databaseSkills.String, Hobbies: databaseHobbies.String}

	stdprofiletpl := stdProfiletpl{Student: student, Username: session.Username, IsOnline: IsOnline}

	htmltpl, err := template.ParseFiles(PPath+"/views/stdProfile.html", navbartpl, bootstraptpl)
	if ErrorHandler(w, err, TemplatePraseError, 500) {
		http.Redirect(w, r, "/stdMessage?mtype=100", 301)
	}

	err = htmltpl.Execute(w, stdprofiletpl)
	if ErrorHandler(w, err, TemplateExecutionError, 500) {
		http.Redirect(w, r, "/stdMessage?mtype=100", 301)
	}
}
