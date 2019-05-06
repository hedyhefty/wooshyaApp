package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"wooshyaApp/Models"
)

type displayJobHdr struct {
	Username string
	IsOnline bool
	Jobs     []Models.JobModel
}

func StdSearchResultPage(w http.ResponseWriter, r *http.Request) {
	//prase http request
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//query url
	v := r.URL.Query()
	keywords := v.Get("keywords")
	search_type := v.Get("searchtype")
	fmt.Println("keyword is ", keywords)
	fmt.Println("search type is ", search_type)

	keywords = strings.ToUpper(keywords)
	//end

	//checkOnline
	IsOnline, session := CheckLogin(Student, r)
	//end

	var displayhdr displayJobHdr
	displayhdr.IsOnline = IsOnline
	if IsOnline {
		displayhdr.Username = session.Username
	}

	stflag, _ := strconv.Atoi(search_type)

	switch stflag {
	case 1:
		rows, err := DB.Query("select id from cpyusers where upper(companyname) like ?", "%"+keywords+"%")
		if ErrorHandler(w, err, "case1 search error.", 500) {
			return
		}

		var idhandle []int
		for rows.Next() {
			var cpy_id int
			err = rows.Scan(&cpy_id)
			if ErrorHandler(w, err, "Scan error.", 500) {
				return
			}

			idhandle = append(idhandle, cpy_id)
		}

		for _, cpyid := range idhandle {
			jrows, err := DB.Query("select jid,jtitle from jobs where cpyid = ?", cpyid)
			if ErrorHandler(w, err, "case1 job query error.", 500) {
				return
			}

			for jrows.Next() {
				job := Models.JobModel{Cpyid: cpyid}
				err = jrows.Scan(&job.Jid, &job.Title)
				if ErrorHandler(w, err, "case1 job scan error", 500) {
					return
				}

				job.JobURL = "/stdSearchResultPage/viewResult?jid=" + strconv.Itoa(job.Jid)

				displayhdr.Jobs = append(displayhdr.Jobs, job)
			}
		}

	case 2:

		rows, err := DB.Query("select jid,jtitle from jobs where upper(jtitle) like ?", "%"+keywords+"%")
		if ErrorHandler(w, err, "case2 query error.", 500) {
			return
		}

		for rows.Next() {
			var job Models.JobModel
			err = rows.Scan(&job.Jid, &job.Title)
			if ErrorHandler(w, err, "case2 scan error.", 500) {
				return
			}

			job.JobURL = "/stdSearchResultPage/viewResult?jid=" + strconv.Itoa(job.Jid)

			displayhdr.Jobs = append(displayhdr.Jobs, job)
		}

	case 3:
		rows, err := DB.Query("select id from cpyusers where upper(category) like ?", "%"+keywords+"%")
		if ErrorHandler(w, err, "case1 search error.", 500) {
			return
		}

		var idhandle []int
		for rows.Next() {
			var cpy_id int
			err = rows.Scan(&cpy_id)
			if ErrorHandler(w, err, "Scan error.", 500) {
				return
			}

			idhandle = append(idhandle, cpy_id)
		}

		for _, cpyid := range idhandle {
			jrows, err := DB.Query("select jid,jtitle from jobs where cpyid = ?", cpyid)
			if ErrorHandler(w, err, "case1 job query error.", 500) {
				return
			}

			for jrows.Next() {
				job := Models.JobModel{Cpyid: cpyid}
				err = jrows.Scan(&job.Jid, &job.Title)
				if ErrorHandler(w, err, "case1 job scan error", 500) {
					return
				}

				job.JobURL = "/stdSearchResultPage/viewResult?jid=" + strconv.Itoa(job.Jid)

				displayhdr.Jobs = append(displayhdr.Jobs, job)
			}
		}
	}
	//end switch

	htmlhdr, err := template.ParseFiles(PPath+"/views/stdSearchResultPage.html", PPath+"/views/navbartpl.html", PPath+"/views/bootstrapHeader.html")

	if ErrorHandler(w, err, "html prase error.", 500) {
		return
	}

	err = htmlhdr.Execute(w, displayhdr)
	if ErrorHandler(w, err, "execute error.", 500) {
		return
	}

}
