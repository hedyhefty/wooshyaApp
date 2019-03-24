package routerHandlers

import (
	"database/sql"
	"os"
)

var DB *sql.DB
var PPath string

func init() {
	var err error
	PPath, err = os.Getwd()
	if err != nil {
		panic(err.Error())
	}
}
