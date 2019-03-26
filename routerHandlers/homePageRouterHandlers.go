package routerHandlers

import (
	"fmt"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	err := DB.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		println("I know u DB.")
	}
	fmt.Println("PPath: ", PPath)
	fmt.Println(PPath + "/views/index.html")
	http.ServeFile(res, req, PPath+"/views/index.html")
}
