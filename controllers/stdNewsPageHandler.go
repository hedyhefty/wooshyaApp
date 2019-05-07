package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

func StdNewsPage(w http.ResponseWriter, r *http.Request) {
	//prase header
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//check user
	IsOnline, session := CheckLogin(Student, r)
	//end check

	//query db to get news title and id
	rows, err := DB.Query("select news_id, news_title from news")

	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	var displayhdr newsListdisplayHdr

	if IsOnline {
		displayhdr.Username = session.Username
		displayhdr.IsOnline = true
	} else {
		displayhdr.IsOnline = false
	}

	for rows.Next() {
		var news_idHdr int
		var news_titleHdr string
		err = rows.Scan(&news_idHdr, &news_titleHdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("news id: ", news_idHdr)
		fmt.Println("news title:", news_titleHdr)

		news_url := "/viewNews?nid=" + strconv.Itoa(news_idHdr)

		newsHdr := Models.NewsModel{NewsURL: news_url, NewsTitle: news_titleHdr}

		displayhdr.News = append(displayhdr.News, newsHdr)
	}

	htmltpl, err := template.ParseFiles(PPath+"/views/stdNewsPage.html", PPath+"/views/navbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		http.Error(w, "Template Prase Error", http.StatusInternalServerError)
		return
	}

	err = htmltpl.Execute(w, displayhdr)
	if err != nil {
		http.Error(w, "Template Execute Error", http.StatusInternalServerError)
	}

}
