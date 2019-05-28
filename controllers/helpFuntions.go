package controllers

import (
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"regexp"
)

const (
	GetIDError = "Get Id Error."

	QueryError = "Query Error."
	ScanError  = "Scan Error."

	TemplatePraseError     = "Template Prase Error."
	TemplateExecutionError = "Template Execution Error."
)

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

func ErrorHandler(w http.ResponseWriter, err error, errorString string, errCode int) bool {
	if err != nil {
		http.Error(w, errorString, errCode)
		fmt.Println(err.Error())
		return true
	}
	return false
}

func GetFromValue(r *http.Request, s string) string {
	return template.HTMLEscapeString(r.FormValue(s))
}

func CheckUserName(username string) (bool, error) {
	matched, err := regexp.MatchString(`^[0-9A-Za-z]+$`, username)
	if len(username) < 6 {
		matched = false
	}

	return matched, err
}

func CheckMailAddress(mailaddress string) (bool, error) {
	matched, err := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, mailaddress)

	return matched, err
}
