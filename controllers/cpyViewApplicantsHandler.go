package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type applicantstplHdr struct {
	Username   string
	applicants []Models.ApplicationModel
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
		return
	}

	tplhdr := applicantstplHdr{Username: session.Username}

	for rows.Next() {
		application := Models.ApplicationModel{Jid: jid}
		err = rows.Scan(&application.Stdid, &application.ApplyDate)
		if ErrorHandler(w, err, ScanError, 500) {
			return
		}
		tplhdr.applicants = append(tplhdr.applicants, application)
	}

	//TODO: template prase and execute, write html file.
}
