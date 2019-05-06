package controllers

import (
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

	username := r.FormValue("username")
	mailaddress := r.FormValue("mailaddress")
	newpassword := r.FormValue("newpassword")

	var databaseMailadd string

	err := DB.QueryRow("SELECT mailaddress FROM cpyusers WHERE username = ?", username).Scan(&databaseMailadd)
	if err != nil {
		fmt.Println("query mail error.")
		return
	}

	if mailaddress == databaseMailadd {
		passwordUPdateHandler, err := DB.Prepare("UPDATE cpyusers SET password = ? WHERE username = ?")
		if err != nil {
			fmt.Println("query failed.")
			return
		}
		hashedpass, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("generate password failed.")
			return
		}
		_, err = passwordUPdateHandler.Exec(hashedpass, username)
		if err != nil {
			fmt.Println("pass word update error.")
			return
		}
		fmt.Println("pass word update succeed.")
		http.Redirect(w, r, "/cpyIndex", http.StatusSeeOther)

	} else {
		fmt.Println("wrong mail address.")
		http.Redirect(w, r, "/cpyForgotPass", http.StatusSeeOther)
	}

}
