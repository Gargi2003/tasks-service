package service

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	utils "tasks/common"
	"time"
)

type response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`

	IssueType   string `json:"issue_type"`
	Assignee    string `json:"assignee"`
	Sprint      int    `json:"sprint_id"`
	ProjectId   int    `json:"project_id"`
	StoryPoints int    `json:"points"`
	Reporter    string `json:"reporter"`
	Comments    string `json:"comments"`
}

func SendEmail(req interface{}) {

	requestJSON, err := json.Marshal(req)
	if err != nil {
		utils.Logger.Err(err).Msg("Error converting req object to JSON")
		return
	}

	request := CreateRequest{}
	if err := json.Unmarshal(requestJSON, &request); err != nil {
		utils.Logger.Err(err).Msg("Error unmarshaling req object")
		return
	}

	from := "notification.taskmanager@gmail.com"
	password := "lpyrceoighokxnfs"
	toList := []string{"gargibanerjee49@gmail.com"}
	host := "smtp.gmail.com"
	port := "587"

	subject := "Subject: [TASK NINJA] Updates for " + request.IssueType + ": " + request.Title + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<!DOCTYPE html><html><body><p>There was <strong>1 update</strong></p><p>This issue is now <strong>assigned to you</strong></p><div style=\"border: 1px solid black; padding: 10px;\"><p>DUMMY  " + strconv.Itoa(request.ProjectId) + " / " + request.Status + "</p><h3><strong>" + request.Title + "</strong></h3><p><a href=\"localhost:4200\">View issue</a></p><p>Changes by <strong>" + request.Reporter + "</strong> on " + time.Now().Format("2006-01-02 15:04:05") + "</p><p style=\"background-color: #AEFF77;width:150px;\">Assignee: " + request.Assignee + "</p></div></body></html>"

	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", from, password, host)
	err1 := smtp.SendMail(host+":"+port, auth, from, toList, msg)
	if err1 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Email sent successfully!")
}
