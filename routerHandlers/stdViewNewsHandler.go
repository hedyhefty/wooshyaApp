package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

func StdViewNews(w http.ResponseWriter, r *http.Request) {
	//prase header
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//check user
	IsOnline, session := CheckLogin(Student, r)
	//end check

	paras := r.URL.Query()
	nid_str := paras.Get("nid")

	nid, err := strconv.Atoi(nid_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	news := Models.NewsModel{NewsID: nid}

	err = DB.QueryRow("select cpy_id,news_title,news_content,release_date from news where news_id = ?", nid).Scan(&news.CpyID, &news.NewsTitle, &news.NewsContent, &news.ReleaseDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = DB.QueryRow("select companyname from cpyusers where id = ?", news.CpyID).Scan(&news.CpyName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var displayhdr displayNewsHdr

	if IsOnline {
		displayhdr = displayNewsHdr{Username: session.Username, IsOnline: IsOnline, News: news}
	} else {
		displayhdr = displayNewsHdr{ IsOnline: IsOnline, News: news}
	}

	htmltpl, err := template.ParseFiles(PPath+"/views/stdViewNews.html", PPath+"/views/navbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		http.Error(w, "Template Prase Error", http.StatusInternalServerError)
		return
	}

	err = htmltpl.Execute(w, displayhdr)
	if err != nil {
		http.Error(w, "Template Excute Error", http.StatusInternalServerError)
	}

}
