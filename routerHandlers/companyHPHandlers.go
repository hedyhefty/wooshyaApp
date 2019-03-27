package routerHandlers

import (
	"net/http"
)

func CompanyHomePage(res http.ResponseWriter, req *http.Request) {

	// fmt.Println("PPath: ", PPath)
	// fmt.Println(PPath + "/views/index.html")
	http.ServeFile(res, req, PPath+"/views/companyIndex.html")
}
