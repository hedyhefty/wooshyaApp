package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"wooshyaApp/Models"
)

type displayResultHdr struct {
	IsOnline  bool
	IsApplied bool
	Username  string
	Job       Models.JobModel
}

func StdViewResult(w http.ResponseWriter, r *http.Request) {
	//prase header
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//check user
	IsOnline, session := CheckLogin(Student, r)
	//end check

	//initialize display handler.
	displayResult := displayResultHdr{IsOnline: IsOnline}
	if IsOnline {
		displayResult.Username = session.Username
	}
	//end

	paras := r.URL.Query()
	jid_str := paras.Get("jid")

	jid, err := strconv.Atoi(jid_str)
	if ErrorHandler(w, err, "strconv failed.", 500) {
		return
	}

	if r.Method == "GET" {
		job := Models.JobModel{Jid: jid}

		jobRow, err := DB.Query("select cpyid,jtitle,jdescribe,jsalary,jlocation,jotherdetails,startdate,deadline from jobs where jid=?", jid)
		if ErrorHandler(w, err, "query failed.", 500) {
			return
		}

		if jobRow.Next() {
			err = jobRow.Scan(&job.Cpyid, &job.Title, &job.Describe, &job.Salary, &job.Location, &job.OtherDetails, &job.StartDate, &job.Deadline)
			if ErrorHandler(w, err, "scan error", 500) {
				fmt.Println(err)
				return
			}
		}

		//find cpyname
		err = DB.QueryRow("select companyname from cpyusers where id=?", job.Cpyid).Scan(&job.CpyName)
		if ErrorHandler(w, err, "cpyname query failed.", 500) {
			return
		}

		//check if is applied
		if IsOnline {
			stdid, err := GetID(session)
			if ErrorHandler(w, err, "get id failed.", 500) {
				return
			}
			var checkhold int
			err = DB.QueryRow("select jid from application where jid=?, stdid=?", jid, stdid).Scan(&checkhold)
			if err == sql.ErrNoRows {
				displayResult.IsApplied = false
			} else {
				displayResult.IsApplied = true
			}
		}

		displayResult.Job = job

		htmlhdr, err := template.ParseFiles(PPath+"/views/stdViewResult.html", PPath+"/views/bootstrapHeader.html", PPath+"/views/navbartpl.html")
		err = htmlhdr.Execute(w, displayResult)
		if ErrorHandler(w, err, "execute error.", 500) {
			return
		}
	}

	if r.Method == "POST" {
		stdid, err := GetID(session)
		fmt.Println(stdid)
		if ErrorHandler(w, err, "get id failed", 500) {
			return
		}
		application := Models.ApplicationModel{Jid: jid, Stdid: stdid}
		applydate := time.Now().Local()
		_, err = DB.Exec("insert into application(jid,stdid,applydate) values (?,?,?)", application.Jid, application.Stdid, applydate)
		if ErrorHandler(w, err, "Insert failed", 500) {
			return
		}

		fmt.Println("applied.")
		http.Redirect(w, r, "/stdSearchResultPage/viewResult", 303)

	}

}
