package routerHandlers

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

func CpySignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		csuptpl, err := template.ParseFiles(PPath+"/views/cpySignUp.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			panic(err.Error())
			return
		}

		err = csuptpl.Execute(w, nil)
		if err != nil {
			panic(err.Error())
			return
		}
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	mailaddress := r.FormValue("mailaddress")
	companyname := r.FormValue("companyname")
	telephonenumber := r.FormValue("telephonenumber")
	lastlogindate := time.Now().Local()

	var cpyuser string
	err := DB.QueryRow("SELECT username FROM cpyusers WHERE username = ?", username).Scan(&cpyuser)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error, unable to create account", 500)
			return
		}

		_, err = DB.Exec("INSERT INTO cpyusers(username,password,mailaddress,companyname,telephonenumber,lastlogindate) VALUES(?,?,?,?,?,?)", username, hashedPassword, mailaddress, companyname, telephonenumber, lastlogindate)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to create account after db.exec", 500)
			return
		}

		fmt.Println("user created.")
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return

	case err != nil:
		http.Error(w, "Server error, unable to create account", 500)
		return

	default:
		http.Redirect(w, r, "/cpySignUp", http.StatusSeeOther)
	}

}
