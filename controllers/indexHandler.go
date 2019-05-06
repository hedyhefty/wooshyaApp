package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Navtpl struct {
	Username string
	IsOnline bool
}

func StdIndex(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("PPath: ", PPath)
	//fmt.Println(PPath + "/views/index.html")

	//deal with the request to favcion, which missing until now.
	if r.URL.Path == "/favicon.ico" {
		return
	}

	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)

	IsOnline, session := CheckLogin(Student, r)

	if r.Method == "GET" {
		navtpl := Navtpl{IsOnline: IsOnline}
		if IsOnline {
			navtpl.Username = session.Username
		}

		htmlHdr, err := template.ParseFiles(PPath+"/views/index.html", navbartpl, bootstraptpl)
		if err != nil {
			http.Error(w, "Prase html failed.", http.StatusInternalServerError)
			return
		}

		err = htmlHdr.Execute(w, navtpl)
		if err != nil {
			http.Error(w, "html template exection failed.", http.StatusInternalServerError)
		}

		return
	}

	if r.Method == "POST" {
		keywords := r.FormValue("keywords")
		searchtype := r.FormValue("search_type")

		res_url := "/stdSearchResultPage?keywords=" + keywords + "&searchtype=" + searchtype

		http.Redirect(w, r, res_url, 303)
		return
	}

}


