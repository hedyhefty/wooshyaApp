package routerHandlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type displayHdr struct {
	Username  string
	NewsTitle []string
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

	rows, err := DB.Query("select news_title from news where cpy_id = ?", cpyid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var displaytplhdr displayHdr
	displaytplhdr.Username = session.Username

	for rows.Next() {
		var stringHandle string
		err = rows.Scan(&stringHandle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("news title:", stringHandle)
		displaytplhdr.NewsTitle = append(displaytplhdr.NewsTitle, stringHandle)
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
