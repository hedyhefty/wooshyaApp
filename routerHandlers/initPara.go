package routerHandlers

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB
var PPath string
var SessionMap map[string]*Session

type Session struct {
	Username   string
	Connection bool
}

func (session *Session) StartSession() {
	fmt.Printf("Sesson %s started.\n", (*session).Username)
}

func init() {
	var err error
	PPath, err = os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	SessionMap = make(map[string]*Session)
}
