package controllers

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

func CpyForgotPass(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		CFPtplHandler, err := template.ParseFiles(PPath+"/views/cpyForgotPass.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			fmt.Println("cfp template parse failed.")
			return
		}

		err = CFPtplHandler.Execute(w, nil)
		if err != nil {
			fmt.Println("cfp excute error.")
			fmt.Println(err)
			return
		}

		return
	}

	username := GetFromValue(r, "username")
	mailaddress := GetFromValue(r, "mailaddress")
	newpassword := r.FormValue("newpassword")
	if len(newpassword) < 6 {
		http.Redirect(w, r, "/cpyMessage?mtype=13", 301)
	}

	var databaseMailadd string

	err := DB.QueryRow("SELECT mailaddress FROM cpyusers WHERE username = ?", username).Scan(&databaseMailadd)

	if err == sql.ErrNoRows {
		http.Redirect(w, r, "/cpyMessage?mtype=12", 301)
	}

	if err != nil {
		//fmt.Println("query mail error.")
		http.Redirect(w, r, "/cpyMessage?mtype=9", 301)
		return
	}

	if mailaddress == databaseMailadd {
		passwordUPdateHandler, err := DB.Prepare("UPDATE cpyusers SET password = ? WHERE username = ?")
		if err != nil {
			//fmt.Println("query failed.")
			http.Redirect(w, r, "/cpyMessage?mtype=9", 301)
			return
		}
		hashedpass, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
		if err != nil {
			//fmt.Println("generate password failed.")
			http.Redirect(w, r, "/cpyMessage?mtype=9", 301)
			return
		}
		_, err = passwordUPdateHandler.Exec(hashedpass, username)
		if err != nil {
			//fmt.Println("pass word update error.")
			http.Redirect(w, r, "/cpyMessage?mtype=9", 301)
			return
		}
		//fmt.Println("pass word update succeed.")
		//http.Redirect(w, r, "/cpyIndex", http.StatusSeeOther)
		http.Redirect(w, r, "/cpyMessage?mtype=10", 301)

	} else {
		//fmt.Println("wrong mail address.")
		//http.Redirect(w, r, "/cpyForgotPass", http.StatusSeeOther)
		http.Redirect(w, r, "/cpyMessage?mtype=11", 301)
	}

}
