package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type displayNewsHdr struct {
	Username string
	News     NewsModel
}

func CpyViewNews(w http.ResponseWriter, r *http.Request) {
	//prase header
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//check user
	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}
	//end

	paras := r.URL.Query()
	nid_str := paras.Get("nid")

	nid, err := strconv.Atoi(nid_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	news := NewsModel{NewsID: nid}

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

	newsDisplayHdr := displayNewsHdr{session.Username, news}

	htmlTemplate, err := template.ParseFiles(PPath+"/views/cpyViewNews.html", PPath+"/views/bootstrapHeader.html", PPath+"/views/hnavbartpl.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = htmlTemplate.Execute(w, newsDisplayHdr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
