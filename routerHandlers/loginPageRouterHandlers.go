package routerHandlers

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func LoginPage(res http.ResponseWriter, req *http.Request) {
	if (*req).Method != "POST" {
		http.ServeFile(res, req, PPath+"/views/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	fmt.Printf("password: %s", password)

	var databaseUsername string
	var databasePassword string

	err := DB.QueryRow("SELECT username,password FROM stdusers where username = ?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		// todo: add a info box including "Username not existed"
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword),
		[]byte(password))
	if err != nil {
		// todo: add a info box including "Wrong password"
		http.Redirect(res, req, "/login", 301)
	}

	updateDatehandler, err := DB.Prepare("UPDATE stdusers SET lastlogindate = ? WHERE username = ?")
	if err != nil {
		// todo: add a info box to inform loginDate error
		http.Redirect(res, req, "/login", 301)
		return
	}

	logindate := time.Now().Local()
	_, err = updateDatehandler.Exec(logindate, username)
	if err != nil {
		http.Redirect(res, req, "/login", 301)
	}

	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 7)
	cookie := http.Cookie{Name: "username", Value: username, Expires: expiration}
	fmt.Println("Setting cookies... cookie: ", cookie)


	http.SetCookie(res, &cookie)

	session := Session{Username:username,Connection:true}
	SessionMap[username] = &session
	session.StartSession()

	res.Write([]byte("hello " + databaseUsername))
}
