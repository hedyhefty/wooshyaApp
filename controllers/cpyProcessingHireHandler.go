package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type cpyDisplayJob struct {
	Username string
	Jobs     []Models.JobModel
}

func CpyProcessingHire(w http.ResponseWriter, r *http.Request) {
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

	cpyid, err := GetID(session)
	if ErrorHandler(w, err, "get id failed.", 500) {
		return
	}

	//Query job belong to the session id.
	rows, err := DB.Query("select jid,jtitle from jobs where cpyid=?", cpyid)
	if ErrorHandler(w, err, "job query failed.", 500) {
		return
	}

	displayhdr := cpyDisplayJob{Username: session.Username}
	fmt.Println(displayhdr.Username)

	for rows.Next() {
		var job Models.JobModel
		err = rows.Scan(&job.Jid, &job.Title)
		job.JobURL = "/cpyIndex/processingHire/viewHire?jid=" + strconv.Itoa(job.Jid)
		displayhdr.Jobs = append(displayhdr.Jobs, job)
	}

	htmltemplate, err := template.ParseFiles(PPath+"/views/cpyProcessingHire.html", hnavbartpl, bootstraptpl)
	if ErrorHandler(w, err, "prase template failed.", 500) {
		return
	}

	err = htmltemplate.Execute(w, displayhdr)
	if ErrorHandler(w, err, "execute template failed.", 500) {
		panic(err.Error())
		return
	}
}
