package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

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
