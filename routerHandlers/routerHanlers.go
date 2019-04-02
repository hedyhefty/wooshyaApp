package routerHandlers

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// )

// func HomePage(res http.ResponseWriter, req *http.Request) {
// 	err := DB.Ping()
// 	if err != nil {
// 		panic(err.Error())
// 	} else {
// 		println("I know u DB.")
// 	}
// 	fmt.Println("PPath: ", PPath)
// 	fmt.Println(PPath + "/views/index.html")
// 	http.ServeFile(res, req, PPath+"/views/index.html")
// }

// func SignupPage(res http.ResponseWriter, req *http.Request) {
// 	if (*req).Method != "POST" {
// 		http.ServeFile(res, req, PPath+"/views/stdSignUp.html")
// 		return
// 	}

// 	username := req.FormValue("username")
// 	password := req.FormValue("password")
// 	mailaddress := req.FormValue("mailaddress")
// 	collegename := req.FormValue("collegename")
// 	degree := req.FormValue("degree")
// 	department := req.FormValue("department")
// 	major := req.FormValue("major")
// 	graduatedate := req.FormValue("graduatedate")
// 	lastlogindate := time.Now().Local()

// 	var stduser string
// 	err := DB.QueryRow("SELECT username FROM stdusers WHERE username = ?", username).Scan(&stduser)
// 	fmt.Printf("Query error type: %s\n", err)

// 	switch {
// 	case err == sql.ErrNoRows:
// 		// Username not exists
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 		if err != nil {
// 			http.Error(res, "Server error, unable to create account", 500)
// 			return
// 		}

// 		_, err = DB.Exec("INSERT INTO stdusers(username,password,mailaddress,collegename,degree,department,major,graduatedate,lastlogindate) VALUES(?,?,?,?,?,?,?,?,?)",
// 			username, hashedPassword, mailaddress, collegename, degree, department, major, graduatedate, lastlogindate)

// 		if err != nil {
// 			http.Error(res, "Server error, unable to create account", 500)
// 			return
// 		}
// 		res.Write([]byte("User Created!"))
// 		http.Redirect(res, req, "/", 200)
// 		// return

// 	case err != nil:
// 		// Other error
// 		http.Error(res, "Server error, unable to create account", 500)
// 		return

// 	default:
// 		// Username already exists
// 		http.Redirect(res, req, "/stdSignUp", 301)
// 	}
// }

// func LoginPage(res http.ResponseWriter, req *http.Request) {
// 	if (*req).Method != "POST" {
// 		http.ServeFile(res, req, PPath+"/views/stdLogin.html")
// 		return
// 	}

// 	username := req.FormValue("username")
// 	password := req.FormValue("password")

// 	var databaseUsername string
// 	var databasePassword string

// 	err := DB.QueryRow("SELECT username,password FROM stdusers where username = ?", username).Scan(&databaseUsername, &databasePassword)

// 	if err != nil {
// 		http.Redirect(res, req, "/stdLogin", 301)
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
// 	if err != nil {
// 		http.Redirect(res, req, "/stdLogin", 301)
// 	}

// 	updateDatehandler, err := DB.Prepare("UPDATE stdusers SET lastlogindate = ? WHERE username = ?")
// 	if err != nil {
// 		http.Redirect(res, req, "/stdLogin", 301)
// 		return
// 	}

// 	logindate := time.Now().Local()
// 	_, err = updateDatehandler.Exec(logindate, username)
// 	if err != nil {
// 		http.Redirect(res, req, "/stdLogin", 301)
// 	}

// 	res.Write([]byte("hello " + databaseUsername))
// }
