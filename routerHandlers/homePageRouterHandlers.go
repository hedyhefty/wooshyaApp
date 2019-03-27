package routerHandlers

import (
	"fmt"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {

	fmt.Println("PPath: ", PPath)
	fmt.Println(PPath + "/views/index.html")
	http.ServeFile(res, req, PPath+"/views/index.html")
}
