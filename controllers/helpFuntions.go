package controllers

import (
	"github.com/pkg/errors"
	"net/http"
)

const(
	GetIDError = "Get Id Error."

	QueryError = "Query Error."
	ScanError = "Scan Error."

	TemplatePraseError = "Template Prase Error."
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
		return true
	}
	return false
}