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

func listenAndServerByMe(addr string, handler http.Handler) error {
	my_server := &http.Server{Addr: addr, Handler: handler}
	fmt.Printf("Server started, listen on port %s\n", addr)
	return my_server.ListenAndServe()
}

func main() {
	//close DB
	defer DB.Close()

	//routers

	http.HandleFunc("/", routerHandlers.StdIndex)
	http.HandleFunc("/stdLogin", routerHandlers.StdLogin)
	http.HandleFunc("/stdSignUp", routerHandlers.StdSignUp)
	http.HandleFunc("/stdLogOut",routerHandlers.StdLogOut)
	http.HandleFunc("/cpyIndex", routerHandlers.CpyIndex)
	http.HandleFunc("/cpyLogin", routerHandlers.CpyLogin)
	http.HandleFunc("/cpySignUp", routerHandlers.CpySignUp)
	http.HandleFunc("/cpyLogOut",routerHandlers.CpyLogOut)
	http.HandleFunc("/cpyForgotPass", routerHandlers.CpyForgotPass)
	lerr := listenAndServerByMe(":8080", nil)
	if lerr != nil {
		panic(lerr.Error())
	}
}
