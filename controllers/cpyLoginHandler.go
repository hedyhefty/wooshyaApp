package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CpyLogin(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {

		//by st
		cpylogintpl, err := template.ParseFiles(PPath+"/views/cpyLogin.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			panic(err.Error())
			return
		}

		err = cpylogintpl.Execute(res, nil)
		if err != nil {
			panic(err.Error())
			return
		}
		//

		//http.ServeFile(res, req, PPath+"/views/cpyLogin.html")
		return
	}

	username := GetFromValue(req, "username")
	fmt.Println(username)
	password := GetFromValue(req, "password")

	var databasePassword string

	err := DB.QueryRow("select password from cpyusers where username = ?", username).Scan(&databasePassword)

	if err == sql.ErrNoRows {
		http.Redirect(res, req, "/cpyMessage?mtype=7", 301)
	}

	if err != nil {
		//http.Redirect(res, req, "/cpyLogin", 301)
		http.Redirect(res, req, "/cpyMessage?mtype=6", 301)
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		//fmt.Println("wrong password.")
		http.Redirect(res, req, "/cpyMessage?mtype=8", 301)
		return
	}

	updateDatehandler, err := DB.Prepare("UPDATE cpyusers SET lastlogindate = ? WHERE username = ?")
	if err != nil {
		http.Redirect(res, req, "/cpyMessage?mtype=6", 301)
		return
	}

	logindate := time.Now().Local()
	_, err = updateDatehandler.Exec(logindate, username)
	if err != nil {
		http.Redirect(res, req, "/cpyMessage?mtype=6", 301)
		return
	}

	session := Session{Username: username, Connection: true, SessionType: Company}
	session.StartSession()
	go session.EndSession()

	cookies := session.SetCookies()
	http.SetCookie(res, &cookies)

	http.Redirect(res, req, "/cpyIndex", http.StatusSeeOther)
}
