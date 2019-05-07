package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type displayhire struct {
	Username string
	Job      Models.JobModel
}

func CpyViewHire(w http.ResponseWriter, r *http.Request) {
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
	jid_str := paras.Get("jid")

	jid, err := strconv.Atoi(jid_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//end

	if r.Method == "POST" {
		newURL := "/cpyIndex/processingHire/viewHire/viewApplicants?jid=" + jid_str
		http.Redirect(w, r, newURL, 303)
		return
	}

	if r.Method == "GET" {
		//query job table for details
		job := Models.JobModel{Jid: jid}
		err = DB.QueryRow("select jtitle,jdescribe,jsalary,jlocation,jotherdetails,releasedate,startdate,deadline from jobs where jid=?", jid).Scan(&job.Title, &job.Describe, &job.Salary, &job.Location, &job.OtherDetails, &job.ReleaseDate, &job.StartDate, &job.Deadline)
		if ErrorHandler(w, err, QueryError, 500) {
			return
		}
		//end

		tplHdr := displayhire{Username: session.Username, Job: job}

		htmlTpl, err := template.ParseFiles(PPath+"/views/cpyViewHire.html", hnavbartpl, bootstraptpl)
		if ErrorHandler(w, err, TemplatePraseError, 500) {
			return
		}

		err = htmlTpl.Execute(w, tplHdr)
		if ErrorHandler(w, err, TemplateExecutionError, 500) {
			return
		}
	}
}
