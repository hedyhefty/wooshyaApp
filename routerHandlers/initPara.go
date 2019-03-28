package routerHandlers

import (
	"database/sql"
	"fmt"
	"os"
	"time"
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

func (session *Session) SetOff() {
	session.Connection = false
	fmt.Printf("Sesson %s was set off, and the state is: %t", (*session).Username, (*session).Connection)
}

func (session *Session) EndSession() {
	var t *time.Timer
	fmt.Printf("EndSession started...\n")
	f := func() {
		session.Connection = false
		fmt.Printf("Session for %s is stopped.\n", session.Username)
		fmt.Printf("C's len: %d\n", len(t.C))
	}

	t = time.AfterFunc(12*time.Minute, f)
	time.Sleep(15 * time.Minute)
}

func init() {
	var err error
	PPath, err = os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	SessionMap = make(map[string]*Session)
}
