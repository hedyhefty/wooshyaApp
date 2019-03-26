package routerHandlers

import (
	"fmt"
	"net/http"
)

func CompanyHomePage(res http.ResponseWriter, req *http.Request) {
	err := DB.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to the database Successfully.")
	}
	// fmt.Println("PPath: ", PPath)
	// fmt.Println(PPath + "/views/index.html")
	http.ServeFile(res, req, PPath+"/views/companyIndex.html")
}
