package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wooshyaApp/Models"
)

type displayHdr struct {
	Username string
	IsOnline bool
	News     []Models.NewsModel
}



func CpyProcessingNews(w http.ResponseWriter, r *http.Request) {
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)

	IsOnline, session := CheckLogin(Company, r)
	if !IsOnline {
		http.Redirect(w, r, "/cpyLogin", http.StatusSeeOther)
		return
	}

	cpyid, err := GetID(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := DB.Query("select news_id, news_title from news where cpy_id = ?", cpyid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var displaytplhdr displayHdr
	displaytplhdr.Username = session.Username

	for rows.Next() {
		var news_idHdr int
		var news_titleHdr string
		err = rows.Scan(&news_idHdr, &news_titleHdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("news id: ", news_idHdr)
		fmt.Println("news title:", news_titleHdr)

		news_url := "/cpyIndex/viewNews?nid=" + strconv.Itoa(news_idHdr)

		newsHdr := Models.NewsModel{NewsURL: news_url, NewsTitle: news_titleHdr}

		displaytplhdr.News = append(displaytplhdr.News, newsHdr)
	}

	htmltemplate, err := template.ParseFiles(PPath+"/views/cpyProcessingNews.html", PPath+"/views/hnavbartpl.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = htmltemplate.Execute(w, displaytplhdr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
