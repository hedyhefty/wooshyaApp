package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"wooshyaApp/routerHandlers"

	_ "github.com/go-sql-driver/mysql"
)

//test contributer
var DB *sql.DB

func init() {
	//start the connection to the DB(mysql).
	DB, err := sql.Open("mysql", "root:Pi3141592653@tcp(127.0.0.1:3306)/wootestdb")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Mysql DB successfully connected.")

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	routerHandlers.DB = DB
}

func main() {
	//close DB
	defer DB.Close()

	mux := http.NewServeMux()

	//routers for student users
	mux.HandleFunc("/", routerHandlers.StdIndex)
	mux.HandleFunc("/stdLogin", routerHandlers.StdLogin)
	mux.HandleFunc("/stdLogOut", routerHandlers.StdLogOut)
	mux.HandleFunc("/stdSignUp", routerHandlers.StdSignUp)
	mux.HandleFunc("/stdForgotPass", routerHandlers.StdForgotPass)

	//routers for company users
	mux.HandleFunc("/cpyIndex",routerHandlers.CpyIndex)
	mux.HandleFunc("/cpyIndex/profile",routerHandlers.CpyProfile)
	mux.HandleFunc("/cpyLogin",routerHandlers.CpyLogin)
	mux.HandleFunc("/cpyLogOut",routerHandlers.CpyLogOut)
	mux.HandleFunc("/cpySignUp",routerHandlers.CpySignUp)
	mux.HandleFunc("/cpyForgotPass",routerHandlers.CpyForgotPass)

	server := &http.Server{Addr:":8080",Handler:mux}
	fmt.Printf("Server started, listen on port %s\n", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
}
