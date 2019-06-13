package controllers

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

	username := GetFromValue(r, "username")
	ok, err := CheckUserName(username)
	if !ok || err != nil {
		//http.Error(w, "Invalid username.", 500)
		http.Redirect(w, r, "/cpyMessage?mtype=3", 301)
		return
	}

	password := r.FormValue("password")
	if len(password) < 6 {
		//http.Error(w, "Invalid password.", 500)
		http.Redirect(w, r, "/cpyMessage?mtype=4", 301)
		return
	}

	mailaddress := GetFromValue(r, "mailaddress")
	ok, err = CheckMailAddress(mailaddress)
	if !ok || err != nil {
		//http.Error(w, "Invalid mailaddress.", 500)
		http.Redirect(w, r, "/cpyMessage?mtype=5", 301)
		return
	}

	companyname := GetFromValue(r, "companyname")
	telephonenumber := GetFromValue(r, "telephonenumber")
	lastlogindate := time.Now().Local()

	//prevent from sign up in the same time.
	cpySignUplocker.Lock()
	defer cpySignUplocker.Unlock()

	var cpyuser string
	err = DB.QueryRow("SELECT username FROM cpyusers WHERE username = ?", username).Scan(&cpyuser)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			//http.Error(w, "Server error, unable to create account", 500)
			http.Redirect(w, r, "/cpyMessage?mtype=2", 301)
			return
		}

		_, err = DB.Exec("INSERT INTO cpyusers(username,password,mailaddress,companyname,telephonenumber,lastlogindate) VALUES(?,?,?,?,?,?)", username, hashedPassword, mailaddress, companyname, telephonenumber, lastlogindate)

		if err != nil {
			fmt.Println(err)
			//http.Error(w, "Server error, unable to create account after db.exec", 500)
			http.Redirect(w, r, "/cpyMessage?mtype=2", 301)
			return
		}

		//fmt.Println("user created.")
		http.Redirect(w, r, "/cpyMessage?mtype=0", http.StatusSeeOther)
		return

	case err != nil:
		//http.Error(w, "Server error, unable to create account", 500)
		http.Redirect(w, r, "/cpyMessage?mtype=2", 301)
		return

	default:
		//http.Redirect(w, r, "/cpySignUp", http.StatusSeeOther)
		http.Redirect(w, r, "/cpyMessage?mtype=1", 301)
	}

}
