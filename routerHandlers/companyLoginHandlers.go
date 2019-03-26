package routerHandlers

import (
	"net/http"

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
		http.Redirect(res, req, "/companyLogin", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword),
		[]byte(password))

}
