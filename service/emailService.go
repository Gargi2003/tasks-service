package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	utils "tasks/common"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	from     = "notification.taskmanager@gmail.com"
	password = "lpyrceoighokxnfs"
	host     = "smtp.gmail.com"
	port     = "587"
)

var toList = []string{"gargibanerjee49@gmail.com"}

type response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Timestamp   time.Time
	IssueType   string `json:"issue_type"`
	Assignee    string `json:"assignee"`
	Sprint      int    `json:"sprint_id"`
	ProjectId   int    `json:"project_id"`
	StoryPoints int    `json:"points"`
	Reporter    string `json:"reporter"`
	Comments    string `json:"comments"`
}

func SendEmailForCreatedIssue(req interface{}) {
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
	convertedTask := &utils.PreviousResponse{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		Assignee:    request.Assignee,
		Reporter:    request.Reporter,
		Comments:    request.Comments,
		Points:      request.StoryPoints,
		ProjectID:   request.ProjectId,
		Sprint:      request.Sprint,
		IssueType:   request.IssueType,
	}
	fileName := "./service/createTemplate.html"
	sendMail(convertedTask, fileName)
	utils.Logger.Info().Msg("Email sent successfully!")
}

func ParseTemplate(fileName string, request interface{}) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		utils.Logger.Err(err).Msg("error while parsing files")
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, request); err != nil {
		utils.Logger.Err(err).Msg("error while applying the template to the data object ")
		return "", err
	}
	body := buf.String()
	return body, nil
}

// Helper function to compare two tasks
func isTaskEqual(task1 *utils.PreviousResponse, task2 *utils.PreviousResponse) bool {
	// Implement the comparison logic based on your requirements
	// Compare the relevant fields to determine if the tasks are the same or different
	return task1.Title == task2.Title &&
		task1.Description == task2.Description &&
		task1.Assignee == task2.Assignee &&
		task1.Reporter == task2.Reporter &&
		task1.Status == task2.Status &&
		task1.IssueType == task2.IssueType &&
		task1.Sprint == task2.Sprint &&
		task1.ProjectID == task2.ProjectID &&
		task1.Points == task2.Points &&
		task1.Comments == task2.Comments
}
func SendEmailForUpdatedIssue(previousTask *utils.PreviousResponse, updatedTask *utils.PreviousResponse) {
	spew.Dump("Previous", previousTask)
	updatedTask.PreviousFields = make(map[string]interface{})
	fileName := "./service/updateTemplate.html"
	updatedTask.Count = 0
	// Compare the previous task and updated task
	if previousTask != nil && updatedTask != nil {
		if !isTaskEqual(previousTask, updatedTask) {
			findCount(previousTask, updatedTask)
			sendMail(updatedTask, fileName)
		}
	}

	utils.Logger.Info().Msg("Email sent successfully!")

}
func sendMail(updatedTask *utils.PreviousResponse, fileName string) {
	subject := "Subject: [TASK NINJA] Updates for " + updatedTask.IssueType + ": " + updatedTask.Title + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	updatedTask.Timestamp = time.Now()
	spew.Dump("Updated", updatedTask)

	body, err := ParseTemplate(fileName, &updatedTask)
	if err != nil {
		utils.Logger.Err(err).Msg("failed to fetch email body")
		return
	}
	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", from, password, host)
	err1 := smtp.SendMail(host+":"+port, auth, from, toList, msg)
	if err1 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func findCount(previousTask *utils.PreviousResponse, updatedTask *utils.PreviousResponse) {
	// Track the updated fields and their values
	if previousTask.Title != updatedTask.Title {
		updatedTask.PreviousFields["Title"] = previousTask.Title
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Description != updatedTask.Description {
		updatedTask.PreviousFields["Description"] = previousTask.Description
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Status != updatedTask.Status {
		updatedTask.PreviousFields["Status"] = previousTask.Status
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Assignee != updatedTask.Assignee {
		updatedTask.PreviousFields["Assignee"] = previousTask.Assignee
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Reporter != updatedTask.Reporter {
		updatedTask.PreviousFields["Reporter"] = previousTask.Reporter
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Comments != updatedTask.Comments {
		updatedTask.PreviousFields["Comments"] = previousTask.Comments
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Points != updatedTask.Points {
		updatedTask.PreviousFields["Points"] = previousTask.Points
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.ProjectID != updatedTask.ProjectID {
		updatedTask.PreviousFields["ProjectID"] = previousTask.ProjectID
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Sprint != updatedTask.Sprint {
		updatedTask.PreviousFields["Sprint"] = previousTask.Sprint
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.IssueType != updatedTask.IssueType {
		updatedTask.PreviousFields["IssueType"] = previousTask.IssueType
		updatedTask.Count = updatedTask.Count + 1
	}
}
