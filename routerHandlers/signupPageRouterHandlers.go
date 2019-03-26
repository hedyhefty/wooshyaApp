package routerHandlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func SignupPage(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {
		http.ServeFile(res, req, PPath+"/views/signup.html")
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
	fmt.Printf("Query error type: %s\n", err)

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
		http.Redirect(res, req, "/signup", 301)
	}
}
