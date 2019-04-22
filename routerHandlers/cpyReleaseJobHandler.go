package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type ReleasetplHdr struct {
	Username string
}

func CpyReleaseJob(w http.ResponseWriter, r *http.Request) {
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)

	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		jtitle := r.FormValue("jtitle")
		jdescribe := r.FormValue("jdescribe")
		jsalary := r.FormValue("jsalary")
		jlocation := r.FormValue("jlocation")
		jotherdetails := r.FormValue("jotherdetails")

		//2019-04-15T01:59
		startdate := r.FormValue("startdate")
		deadline := r.FormValue("deadline")

		startdate = PraseDateTime(startdate)
		deadline = PraseDateTime(deadline)
		releasedate := time.Now().Local()

		fmt.Println(startdate)
		fmt.Println(deadline)

		var cpyid int
		err := DB.QueryRow("SELECT id FROM cpyusers WHERE username = ?", session.Username).Scan(&cpyid)
		if err != nil {
			panic(err.Error())
		}

		_, err = DB.Exec("INSERT INTO jobs(cpyid,jtitle,jdescribe,jsalary,jlocation,jotherdetails,releasedate,startdate,deadline) VALUES(?,?,?,?,?,?,?,?,?)", cpyid, jtitle, jdescribe, jsalary, jlocation, jotherdetails, releasedate, startdate, deadline)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to release job after db.exec", 500)
			return
		}

		fmt.Println("Job released.")
		http.Redirect(w, r, "/cpyIndex", http.StatusSeeOther)
		return
	}

	if r.Method == "GET" {
		reltplHdr := ReleasetplHdr{Username: session.Username}
		releaseJobtpl, err := template.ParseFiles(PPath+"/views/cpyReleaseJob.html", PPath+"/views/hnavbartpl.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			panic(err.Error())
		}

		err = releaseJobtpl.Execute(w, reltplHdr)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
}
