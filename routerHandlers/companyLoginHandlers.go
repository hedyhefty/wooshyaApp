package routerHandlers

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CompanyLogin(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {
		http.ServeFile(res, req, PPath+"/views/companyLogin.html")
		return
	}

	companyName := req.FormValue("companyname")
	password := req.FormValue("password")

	var databaseCompanyName string
	var databasePassword string

	err := DB.QueryRow("select companyname, password from companyuser where companyname = ?", companyName).Scan(&databaseCompanyName, &databasePassword)
	if err != nil {
		// todo: add a info box including "Username not existed"
		http.Redirect(res, req, "/companyLogin", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword),
		[]byte(password))
	if err != nil {
		// todo: add a info box including "Wrong password"
		http.Redirect(res, req, "/companyLogin", 301)
	}

	updateDatehandler, err := DB.Prepare("UPDATE companyuser SET lastlogindate = ? WHERE username = ?")
	if err != nil {
		// todo: add a info box to inform loginDate error
		http.Redirect(res, req, "/companyLogin", 301)
		return
	}

	logindate := time.Now().Local()
	_, err = updateDatehandler.Exec(logindate, companyName)
	if err != nil {
		http.Redirect(res, req, "/companyLogin", 301)
	}

	res.Write([]byte("hello " + databaseCompanyName))
}
