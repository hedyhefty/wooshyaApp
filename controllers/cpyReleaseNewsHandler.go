package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func CpyReleaseNews(w http.ResponseWriter, r *http.Request) {
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)

	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}

	if r.Method == "GET" {
		releaseNewstplHdr := ReleasetplHdr{session.Username}
		releaseNewstpl, err := template.ParseFiles(PPath+"/views/cpyReleaseNews.html", PPath+"/views/hnavbartpl.html", PPath+"/views/bootstrapHeader.html")
		if err != nil {
			http.Error(w, "Prase html error.", http.StatusInternalServerError)
			return
		}

		err = releaseNewstpl.Execute(w, releaseNewstplHdr)
		if err != nil {
			http.Error(w, "Template execute Error.", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == "POST" {
		news_title := r.FormValue("news_title")
		news_content := r.FormValue("news_content")
		release_date := time.Now().Local()
		var cpy_id int
		err := DB.QueryRow("select id from cpyusers where username = ?", session.Username).Scan(&cpy_id)
		if err != nil {
			http.Error(w, "Query error in release news.", http.StatusInternalServerError)
			return
		}

		_, err = DB.Exec("insert into news(cpy_id,news_title,news_content,release_date) values(?,?,?,?)", cpy_id, news_title, news_content, release_date)
		if err != nil {
			http.Error(w, "Insert error in release news.", http.StatusInternalServerError)
			return
		}
		fmt.Println("News released.")
		http.Redirect(w, r, "/cpyIndex", http.StatusSeeOther)
		return
	}

}
