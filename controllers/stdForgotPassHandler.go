package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

func StdForgotPass(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SFPtplHandler, err := template.ParseFiles(PPath+"/views/stdForgotPass.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			fmt.Println("sfp template parse failed.")
			return
		}

		err = SFPtplHandler.Execute(w, nil)
		if err != nil {
			fmt.Println("sfp excute error.")
			fmt.Println(err)
			return
		}

		return
	}

	username := GetFromValue(r, "username")
	mailaddress := GetFromValue(r, "mailaddress")
	newpassword := r.FormValue("newpassword")

	var databaseMailadd string

	err := DB.QueryRow("SELECT mailaddress FROM stdusers WHERE username = ?", username).Scan(&databaseMailadd)
	if err != nil {
		fmt.Println("query mail error.")
		return
	}

	if mailaddress == databaseMailadd {
		passwordUPdateHandler, err := DB.Prepare("UPDATE stdusers SET password = ? WHERE username = ?")
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
		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		fmt.Println("wrong mail address.")
		http.Redirect(w, r, "/stdForgotPass", http.StatusSeeOther)
	}
}
