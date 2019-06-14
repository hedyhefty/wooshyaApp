package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type stdDisplayJob struct {
	Username string
	IsOnline bool
	Jobs     []Models.JobModel
}

func StdViewApplied(w http.ResponseWriter, r *http.Request) {
	//prase header
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	IsOnline, session := CheckLogin(Student, r)
	if !IsOnline {
		http.Redirect(w, r, "/stdMessage?mtype=14", 303)
	}

	stdid, err := GetID(session)
	if ErrorHandler(w, err, QueryError, 500) {
		return
	}

	rows, err := DB.Query("select jid from application where stdid=?", stdid)
	stddisplayjob := stdDisplayJob{Username: session.Username, IsOnline: IsOnline}

	for rows.Next() {
		var jid int
		var jtit string
		err = rows.Scan(&jid)
		if ErrorHandler(w, err, QueryError, 500) {
			return
		}
		row := DB.QueryRow("select jtitle from jobs where jid=?", jid)
		err = row.Scan(&jtit)
		if ErrorHandler(w, err, QueryError, 500) {
			return
		}
		job := Models.JobModel{Title: jtit, JobURL: "/stdSearchResultPage/viewResult?jid=" + strconv.Itoa(jid)}

		stddisplayjob.Jobs = append(stddisplayjob.Jobs, job)
	}

	htmltpl, err := template.ParseFiles(PPath+"/views/stdViewApplied.html", navbartpl, bootstraptpl)
	if ErrorHandler(w, err, TemplatePraseError, 500) {
		return
	}

	err = htmltpl.Execute(w, stddisplayjob)
	if ErrorHandler(w, err, TemplateExecutionError, 500) {
		return
	}
}
