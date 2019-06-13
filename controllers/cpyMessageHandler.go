package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func CpyMessage(w http.ResponseWriter, r *http.Request) {
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
		message.RedirectURL = "/cpyLogin"
	case 1:
		message.MessageContent = "Username have been used."
		message.RedirectURL = "/cpySignUp"
	case 2:
		message.MessageContent = "Some error occur at server."
		message.RedirectURL = "/cpySignUp"
	case 3:
		message.MessageContent = "Invalid username."
		message.RedirectURL = "/cpySignUp"
	case 4:
		message.MessageContent = "Invalid password."
		message.RedirectURL = "/cpySignUp"
	case 5:
		message.MessageContent = "Invalid mail address."
		message.RedirectURL = "/cpySignUp"
	case 6:
		message.MessageContent = "Some error occur at server."
		message.RedirectURL = "/cpyLogin"
	case 7:
		message.MessageContent = "Wrong username."
		message.RedirectURL = "/cpyLogin"
	case 8:
		message.MessageContent = "Wrong password."
		message.RedirectURL = "/cpyLogin"
	case 9:
		message.MessageContent = "Some error occur at server."
		message.RedirectURL = "/cpyForgotPass"
	case 10:
		message.MessageContent = "Password reset successfully."
		message.RedirectURL = "/cpyLogin"
	case 11:
		message.MessageContent = "Wrong mail address."
		message.RedirectURL = "/cpyForgotPass"
	case 12:
		message.MessageContent = "Username not exist."
		message.RedirectURL = "/cpyForgotPass"
	case 13:
		message.MessageContent = "Invalid password."
		message.RedirectURL = "/cpyForgotPass"
	case 14:
		message.MessageContent = "You are offline, please login."
		message.RedirectURL = "/cpyLogin"

	default:
		message.MessageContent = "Default message."
		message.RedirectURL = "/cpyLogin"
	}

	htmltpl, err := template.ParseFiles(PPath+"/views/cpyMessage.html", bootstraptpl)
	if ErrorHandler(w, err, TemplatePraseError, 500) {
		return
	}

	err = htmltpl.Execute(w, message)
	if ErrorHandler(w, err, TemplateExecutionError, 500) {
		return
	}
}
