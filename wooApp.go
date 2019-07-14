package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"wooshyaApp/controllers"

	_ "github.com/go-sql-driver/mysql"
)

//test contributer
var DB *sql.DB

func init() {
	//start the connection to the DB(mysql).
	DB, err := sql.Open("mysql", "[username]:[password]@[dbconnection]/[dbname]")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Mysql DB successfully connected.")

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	controllers.DB = DB
}

//handle err for defer db.close
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB has closed.")
	}
}

func main() {
	//close DB
	defer CloseDB(DB)

	mux := http.NewServeMux()

	mux.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	//routers for student users
	mux.HandleFunc("/", controllers.StdIndex)
	mux.HandleFunc("/stdSearchResultPage", controllers.StdSearchResultPage)
	mux.HandleFunc("/stdSearchResultPage/viewResult", controllers.StdViewResult)
	mux.HandleFunc("/stdNewsPage", controllers.StdNewsPage)
	mux.HandleFunc("/viewNews", controllers.StdViewNews)
	mux.HandleFunc("/stdLogin", controllers.StdLogin)
	mux.HandleFunc("/stdLogOut", controllers.StdLogOut)
	mux.HandleFunc("/stdSignUp", controllers.StdSignUp)
	mux.HandleFunc("/stdForgotPass", controllers.StdForgotPass)
	mux.HandleFunc("/stdMessage", controllers.StdMessage)
	mux.HandleFunc("/stdProfile", controllers.StdProfile)
	mux.HandleFunc("/stdViewApplied", controllers.StdViewApplied)

	//routers for company users
	mux.HandleFunc("/cpyIndex", controllers.CpyIndex)
	mux.HandleFunc("/cpyIndex/profile", controllers.CpyProfile)
	mux.HandleFunc("/cpyIndex/releaseJob", controllers.CpyReleaseJob)
	mux.HandleFunc("/cpyIndex/releaseNews", controllers.CpyReleaseNews)
	mux.HandleFunc("/cpyIndex/processingNews", controllers.CpyProcessingNews)
	mux.HandleFunc("/cpyIndex/viewNews", controllers.CpyViewNews)
	mux.HandleFunc("/cpyIndex/processingHire", controllers.CpyProcessingHire)
	mux.HandleFunc("/cpyIndex/processingHire/viewHire", controllers.CpyViewHire)
	mux.HandleFunc("/cpyIndex/processingHire/viewHire/viewApplicants", controllers.CpyViewApplicants)
	mux.HandleFunc("/cpyIndex/processingHire/viewHire/viewApplicant/applicantProfile", controllers.CpyViewAppProfile)

	mux.HandleFunc("/cpyLogin", controllers.CpyLogin)
	mux.HandleFunc("/cpyLogOut", controllers.CpyLogOut)
	mux.HandleFunc("/cpySignUp", controllers.CpySignUp)
	mux.HandleFunc("/cpyForgotPass", controllers.CpyForgotPass)
	mux.HandleFunc("/cpyMessage", controllers.CpyMessage)

	server := &http.Server{Addr: ":8080", Handler: mux}
	fmt.Printf("Server started, listen on port %s\n", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
}
