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
	ProjectID   int    `json:"project_id"`
	Points      int    `json:"points"`
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
	convertedTask := &utils.Task{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		Assignee:    request.Assignee,
		Reporter:    request.Reporter,
		Comments:    request.Comments,
		Points:      request.Points,
		ProjectID:   request.ProjectID,
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
func isTaskEqual(task1 *utils.Task, task2 *utils.Task) bool {
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

func SendEmailForUpdatedIssue(previousTask *utils.Task, updatedTask *utils.Task) {
	// spew.Dump("Previous", previousTask)
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

func sendMail(updatedTask *utils.Task, fileName string) {
	subject := "Subject: [TASK NINJA] Updates for " + updatedTask.IssueType + ": " + updatedTask.Title + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	updatedTask.Timestamp = time.Now()
	// spew.Dump("Updated", updatedTask)

	body, err := ParseTemplate(fileName, updatedTask)
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

func findCount(previousTask *utils.Task, updatedTask *utils.Task) {
	// Track the updated fields and their values
	if previousTask.Title != updatedTask.Title {
		updatedTask.PreviousFields["Title"] = map[string]interface{}{
			"Previous": previousTask.Title,
			"Updated":  updatedTask.Title,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Description != updatedTask.Description {
		updatedTask.PreviousFields["Description"] = map[string]interface{}{
			"Previous": previousTask.Description,
			"Updated":  updatedTask.Description,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Status != updatedTask.Status {
		updatedTask.PreviousFields["Status"] = map[string]interface{}{
			"Previous": previousTask.Status,
			"Updated":  updatedTask.Status,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Assignee != updatedTask.Assignee {
		updatedTask.PreviousFields["Assignee"] = map[string]interface{}{
			"Previous": previousTask.Assignee,
			"Updated":  updatedTask.Assignee,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Reporter != updatedTask.Reporter {
		updatedTask.PreviousFields["Reporter"] = map[string]interface{}{
			"Previous": previousTask.Reporter,
			"Updated":  updatedTask.Reporter,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Comments != updatedTask.Comments {
		updatedTask.PreviousFields["Comments"] = map[string]interface{}{
			"Previous": previousTask.Comments,
			"Updated":  updatedTask.Comments,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Points != updatedTask.Points {
		updatedTask.PreviousFields["Points"] = map[string]interface{}{
			"Previous": previousTask.Points,
			"Updated":  updatedTask.Points,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.ProjectID != updatedTask.ProjectID {
		updatedTask.PreviousFields["ProjectID"] = map[string]interface{}{
			"Previous": previousTask.ProjectID,
			"Updated":  updatedTask.ProjectID,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.Sprint != updatedTask.Sprint {
		updatedTask.PreviousFields["Sprint"] = map[string]interface{}{
			"Previous": previousTask.Sprint,
			"Updated":  updatedTask.Sprint,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
	if previousTask.IssueType != updatedTask.IssueType {
		updatedTask.PreviousFields["IssueType"] = map[string]interface{}{
			"Previous": previousTask.IssueType,
			"Updated":  updatedTask.IssueType,
		}
		updatedTask.Count = updatedTask.Count + 1
	}
}
