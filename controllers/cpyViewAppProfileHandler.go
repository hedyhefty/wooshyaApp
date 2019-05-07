package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type appProfileHdr struct {
	Username string
	Student  Models.StdUserModel
}

func CpyViewAppProfile(w http.ResponseWriter, r *http.Request) {
	//prase header
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//check user
	IsOnline, session := CheckLogin(Company, r)
	//end check

	//initialize display handler.
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}
	//end

	//prase query parameter
	paras := r.URL.Query()
	sid_str := paras.Get("sid")

	sid, err := strconv.Atoi(sid_str)
	if ErrorHandler(w, err, "ATOI", 500) {
		return
	}
	//end

	stdProfile := Models.StdUserModel{StdID: sid}
	row := DB.QueryRow("select firstname,lastname,collegename,degree,department,major,graduatedate from stdusers where id=?", stdProfile.StdID)
	err = row.Scan(&stdProfile.FirstName, &stdProfile.LastName, &stdProfile.CollegeName, &stdProfile.Degree, &stdProfile.Department, &stdProfile.Major, &stdProfile.GraduateDate)
	if ErrorHandler(w, err, QueryError, 500) {
		return
	}

	displayhdr := appProfileHdr{Username: session.Username, Student: stdProfile}
	htmlTpl, err := template.ParseFiles(PPath+"/views/cpyViewAppProfile.html", hnavbartpl, bootstraptpl)
	if ErrorHandler(w, err, TemplatePraseError, 500) {
		return
	}

	err = htmlTpl.Execute(w, displayhdr)
	if ErrorHandler(w, err, TemplateExecutionError, 500) {
		return
	}

}
