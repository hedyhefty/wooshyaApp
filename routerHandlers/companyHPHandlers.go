package routerHandlers

import (
	"html/template"
	"net/http"
)

func CompanyHomePage(res http.ResponseWriter, req *http.Request) {

	// fmt.Println("PPath: ", PPath)
	// fmt.Println(PPath + "/views/index.html")

	//add by st
	cpyHptpl, err := template.ParseFiles(PPath+"/views/companyIndex.html", PPath+"/views/bootstrapHeader.html")
	if err != nil {
		panic(err.Error())
		return
	}
	err = cpyHptpl.Execute(res, nil)
	if err != nil {
		panic(err.Error())
		return
	}
	//

	//http.ServeFile(res, req, PPath+"/views/companyIndex.html")
}
