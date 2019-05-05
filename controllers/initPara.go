package routerHandlers

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strings"
)

var DB *sql.DB
var PPath string
var SessionMap map[string]*Session


func init() {
	var err error
	PPath, err = os.Getwd()
	if err != nil {
		panic(err.Error())
	}
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

func PraseDateTime(t string) string {
	tmp := []byte(t)
	tmp[10] = ' '
	res := string(tmp)
	return res + ":00"
}

func GetID(session *Session) (int, error) {
	var id int
	if session.SessionType == Student {
		err := DB.QueryRow("SELECT id FROM stdusers WHERE username = ?", session.Username).Scan(&id)
		if err != nil {
			panic(err.Error())
			return -1, errors.New("Unknow error form GetID.")
		}
		return id, nil
	}

	if session.SessionType == Company {
		err := DB.QueryRow("SELECT id FROM cpyusers WHERE username = ?", session.Username).Scan(&id)
		if err != nil {
			panic(err.Error())
			return -1, errors.New("Unknow error form GetID.")
		}
		return id, nil
	}

	return -1, errors.New("Unknow error form GetID.")
}
