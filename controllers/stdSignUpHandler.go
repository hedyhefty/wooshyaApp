package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"
	"wooshyaApp/Models"

	"golang.org/x/crypto/bcrypt"
)

func StdSignUp(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {
		signuptpl, err := template.ParseFiles(PPath+"/views/stdSignUp.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			panic(err.Error())
			return
		}
		err = signuptpl.Execute(res, nil)
		if err != nil {
			panic(err.Error())
			return
		}
		//http.ServeFile(res, req, PPath+"/views/stdSignUp.html")
		return
	}

	var stduser Models.StdUserModel

	stduser.Username = req.FormValue("username")
	stduser.Password = req.FormValue("password")
	stduser.FirstName = req.FormValue("firstname")
	stduser.LastName = req.FormValue("lastname")
	stduser.MailAddress = req.FormValue("mailaddress")
	stduser.CollegeName = req.FormValue("collegename")
	stduser.Degree = req.FormValue("degree")
	stduser.Department = req.FormValue("department")
	stduser.Major = req.FormValue("major")
	stduser.GraduateDate = req.FormValue("graduatedate")

	lastlogindate := time.Now().Local()

	var check_duplicate string
	err := DB.QueryRow("SELECT username FROM stdusers WHERE username = ?", stduser.Username).Scan(&check_duplicate)

	switch {
	case err == sql.ErrNoRows:
		// Username not exists
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(stduser.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create account", 500)
			return
		}

		_, err = DB.Exec("INSERT INTO stdusers(username,password,firstname,lastname,mailaddress,collegename,degree,department,major,graduatedate,lastlogindate) VALUES(?,?,?,?,?,?,?,?,?,?,?)",
			stduser.Username, hashedPassword, stduser.FirstName, stduser.LastName, stduser.MailAddress, stduser.CollegeName, stduser.Degree, stduser.Department, stduser.Major, stduser.GraduateDate, lastlogindate)

		if err != nil {
			http.Error(res, "Server error, unable to create account", 500)
			return
		}
		res.Write([]byte("User Created!"))
		http.Redirect(res, req, "/", 200)
		// return

	case err != nil:
		// Other error
		http.Error(res, "Server error, unable to create account", 500)
		return

	default:
		// Username already exists
		http.Redirect(res, req, "/stdSignUp", 301)
	}
}
