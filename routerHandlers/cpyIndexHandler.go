package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func CpyIndex(w http.ResponseWriter, r *http.Request) {

	//add by st
	fmt.Println("call Chp.")
	cpyHptpl, err := template.ParseFiles(PPath+"/views/cpyIndex.html", PPath+"/views/hnavbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		panic(err.Error())
		return
	}

	var cpyHptplHandler Indextpl

	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}

	cpyHptplHandler.Username = session.Username

	err = cpyHptpl.Execute(w, cpyHptplHandler)
	if err != nil {
		panic(err.Error())
		return
	}
	//

	//http.ServeFile(res, req, PPath+"/views/companyIndex.html")
}
