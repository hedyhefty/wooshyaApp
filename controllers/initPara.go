package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var DB *sql.DB
var PPath string
var navbartpl string
var hnavbartpl string
var bootstraptpl string
var stdSignUplocker sync.Mutex
var cpySignUplocker sync.Mutex

var SessionMap map[string]*Session

func init() {
	var err error
	PPath, err = os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	navbartpl = PPath + "/views/navbartpl.html"
	hnavbartpl = PPath + "/views/hnavbartpl.html"
	bootstraptpl = PPath + "/views/bootstrapHeader.html"
	SessionMap = make(map[string]*Session)
}

//formatRequest generate ascii representation of a request.
func FormatRequest(r *http.Request) string {
	//Create return string.
	var request []string

	//Add request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	//Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	//If this a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	//Return the request as a string
	return strings.Join(request, "\n") + "\n"
}

