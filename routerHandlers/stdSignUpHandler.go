package routerHandlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"

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

	username := req.FormValue("username")
	password := req.FormValue("password")
	mailaddress := req.FormValue("mailaddress")
	collegename := req.FormValue("collegename")
	degree := req.FormValue("degree")
	department := req.FormValue("department")
	major := req.FormValue("major")
	graduatedate := req.FormValue("graduatedate")
	lastlogindate := time.Now().Local()

	var stduser string
	err := DB.QueryRow("SELECT username FROM stdusers WHERE username = ?", username).Scan(&stduser)

	switch {
	case err == sql.ErrNoRows:
		// Username not exists
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create account", 500)
			return
		}

		_, err = DB.Exec("INSERT INTO stdusers(username,password,mailaddress,collegename,degree,department,major,graduatedate,lastlogindate) VALUES(?,?,?,?,?,?,?,?,?)",
			username, hashedPassword, mailaddress, collegename, degree, department, major, graduatedate, lastlogindate)

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
