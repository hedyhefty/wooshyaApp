package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type applicantstplHdr struct {
	Username string
	Students []Models.StdUserModel
}

func CpyViewApplicants(w http.ResponseWriter, r *http.Request) {
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
	if ErrorHandler(w, err, "ATOI", 500) {
		return
	}
	//end

	//query application table

	rows, err := DB.Query("select stdid,applydate from application where jid=?", jid)
	if ErrorHandler(w, err, QueryError, 500) {
		panic(err.Error())
		return
	}

	tplhdr := applicantstplHdr{Username: session.Username}
	var applicants []Models.ApplicationModel

	for rows.Next() {
		application := Models.ApplicationModel{Jid: jid}
		err = rows.Scan(&application.Stdid, &application.ApplyDate)
		if ErrorHandler(w, err, ScanError, 500) {
			return
		}

		applicants = append(applicants, application)

	}

	for _, v := range applicants {
		stdhold := Models.StdUserModel{StdID: v.Stdid}
		err := DB.QueryRow("select firstname,lastname from stdusers where id=?", v.Stdid).Scan(&stdhold.FirstName, &stdhold.LastName)
		if ErrorHandler(w, err, QueryError, 500) {
			panic(err.Error())
			return
		}

		stdhold.ApplyDate = v.ApplyDate

		stdhold.StdURL = "/cpyIndex/processingHire/viewHire/viewApplicant/applicantProfile?sid=" + strconv.Itoa(stdhold.StdID)

		tplhdr.Students = append(tplhdr.Students, stdhold)
	}

	htmltpl, err := template.ParseFiles(PPath+"/views/cpyViewApplicants.html", hnavbartpl, bootstraptpl)
	if ErrorHandler(w, err, TemplatePraseError, 500) {
		return
	}

	err = htmltpl.Execute(w, tplhdr)
	if ErrorHandler(w, err, TemplateExecutionError, 500) {
		panic(err.Error())
		return
	}
}
