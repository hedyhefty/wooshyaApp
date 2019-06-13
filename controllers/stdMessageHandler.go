package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	MessageContent string
	RedirectURL    string
}

func StdMessage(w http.ResponseWriter, r *http.Request) {
	//prase http request
	requestHeader := FormatRequest(r)
	fmt.Printf("\n")
	fmt.Println(requestHeader)
	//end

	//query url
	v := r.URL.Query()
	mtype := v.Get("mtype")
	fmt.Println(mtype)
	//end

	imtype, err := strconv.Atoi(mtype)
	if ErrorHandler(w, err, "Atoi", 500) {
		return
	}

	var message MessageHandler

	switch imtype {
	case 0:
		message.MessageContent = "You have signed up successfully."
		message.RedirectURL = "/stdLogin"
	case 1:
		message.MessageContent = "Username have been used."
		message.RedirectURL = "/stdSignUp"
	case 2:
		message.MessageContent = "Some error occur at server."
		message.RedirectURL = "/stdSignUp"
	case 3:
		message.MessageContent = "Invalid username."
		message.RedirectURL = "/stdSignUp"
	case 4:
		message.MessageContent = "Invalid password."
		message.RedirectURL = "/stdSignUp"
	case 5:
		message.MessageContent = "Invalid mail address."
		message.RedirectURL = "/stdSignUp"
	case 6:
		message.MessageContent = "Some error occur at server."
		message.RedirectURL = "/stdLogin"
	case 7:
		message.MessageContent = "Wrong username."
		message.RedirectURL = "/stdLogin"
	case 8:
		message.MessageContent = "Wrong password."
		message.RedirectURL = "/stdLogin"
	case 9:
		message.MessageContent = "Some error occur at server."
		message.RedirectURL = "/stdForgotPass"
	case 10:
		message.MessageContent = "Password reset successfully."
		message.RedirectURL = "/stdLogin"
	case 11:
		message.MessageContent = "Wrong mail address."
		message.RedirectURL = "/stdForgotPass"
	case 12:
		message.MessageContent = "Username not exist."
		message.RedirectURL = "/stdForgotPass"
	case 13:
		message.MessageContent = "Invalid password."
		message.RedirectURL = "/stdForgotPass"
	case 14:
		message.MessageContent = "You are offline, please login."
		message.RedirectURL = "/stdLogin"

	default:
		message.MessageContent = "Default error message."
		message.RedirectURL = "/"
	}

	htmltpl, err := template.ParseFiles(PPath+"/views/stdMessage.html", bootstraptpl)
	if ErrorHandler(w, err, TemplatePraseError, 500) {
		return
	}

	err = htmltpl.Execute(w, message)
	if ErrorHandler(w, err, TemplateExecutionError, 500) {
		return
	}
}
