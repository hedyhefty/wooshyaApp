package routerHandlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var DB *sql.DB
var PPath string
var SessionMap map[string]*Session

type SessionTypeMask int

const (
	Student = iota
	Company
)

type Session struct {
	Username    string
	Connection  bool
	SessionType SessionTypeMask
	SessionID   string
}

func (session *Session) StartSession() {
	session.MakeSessionID()
	SessionMap[(*session).SessionID] = session
	fmt.Printf("Sesson %s started with Session ID of %s.\n", (*session).Username, (*session).SessionID)
}

func (session *Session) MakeSessionID() {
	sid := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, sid); err != nil {
		(*session).SessionID = ""
	} else {
		(*session).SessionID = base64.URLEncoding.EncodeToString(sid)
	}
}

func (session Session) SetCookies() http.Cookie {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 7)
	cookie := http.Cookie{Name: "SessionID", Value: session.SessionID, Expires: expiration}
	fmt.Println("Setting cookies... cookie: ", cookie)
	return cookie
}

func (session *Session) SetOff() {
	session.Connection = false
	fmt.Printf("Sesson %s was set off, and the state is: %t\n", (*session).Username, (*session).Connection)
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

func CheckLogin(sessionType SessionTypeMask, req *http.Request) (bool, *Session) {
	cookies, err := req.Cookie("SessionID")
	if err != nil {
		fmt.Println("load cookies error: ", err)
		return false, nil
	}

	if SessionMap[cookies.Value] == nil {
		fmt.Println("no relative cookies.")
		return false, nil
	}

	if SessionMap[cookies.Value].SessionType != sessionType {
		fmt.Println("session type missmatch.")
		return false, nil
	}

	if !SessionMap[cookies.Value].Connection {
		fmt.Println("session disconnected.")
		return false, nil
	}

	return true, SessionMap[cookies.Value]
}

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
	return strings.Join(request, "\n")
}
