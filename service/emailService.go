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

type response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type CreateRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Timestamp      time.Time
	Count          int
	PreviousFields map[string]interface{}
	IssueType      string `json:"issue_type"`
	Assignee       string `json:"assignee"`
	Sprint         int    `json:"sprint_id"`
	ProjectId      int    `json:"project_id"`
	StoryPoints    int    `json:"points"`
	Reporter       string `json:"reporter"`
	Comments       string `json:"comments"`
}

func SendEmail(req interface{}, previousTask *utils.PreviousResponse, updatedTask *utils.PreviousResponse) {
	fmt.Println("previous task", previousTask)
	fmt.Println("updated task", updatedTask)
	fmt.Println("req", req)
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
	request = CreateRequest{
		PreviousFields: make(map[string]interface{}),
	}
	request.Count = 0
	// Compare the previous task and updated task
	if previousTask != nil && updatedTask != nil {
		if !isTaskEqual(previousTask, updatedTask) {

			// Track the updated fields and their values
			if previousTask.Title != updatedTask.Title {
				request.PreviousFields["Title"] = previousTask.Title
				request.Count = request.Count + 1
			}
			if previousTask.Description != updatedTask.Description {
				request.PreviousFields["Description"] = previousTask.Description
				request.Count = request.Count + 1
			}
			if previousTask.Status != updatedTask.Status {
				request.PreviousFields["Status"] = previousTask.Status
				request.Count = request.Count + 1
			}
			if previousTask.Assignee != updatedTask.Assignee {
				request.PreviousFields["Assignee"] = previousTask.Assignee
				request.Count = request.Count + 1
			}
			if previousTask.Reporter != updatedTask.Reporter {
				request.PreviousFields["Reporter"] = previousTask.Reporter
				request.Count = request.Count + 1
			}
			if previousTask.Comments != updatedTask.Comments {
				request.PreviousFields["Comments"] = previousTask.Comments
				request.Count = request.Count + 1
			}
			if previousTask.Points != updatedTask.Points {
				request.PreviousFields["Points"] = previousTask.Points
				request.Count = request.Count + 1
			}
			if previousTask.ProjectID != updatedTask.ProjectID {
				request.PreviousFields["ProjectID"] = previousTask.ProjectID
				request.Count = request.Count + 1
			}
			if previousTask.Sprint != updatedTask.Sprint {
				request.PreviousFields["Sprint"] = previousTask.Sprint
				request.Count = request.Count + 1
			}
			if previousTask.IssueType != updatedTask.IssueType {
				request.PreviousFields["IssueType"] = previousTask.IssueType
				request.Count = request.Count + 1
			}

		}
	}
	from := "notification.taskmanager@gmail.com"
	password := "lpyrceoighokxnfs"
	toList := []string{"gargibanerjee49@gmail.com"}
	host := "smtp.gmail.com"
	port := "587"

	subject := "Subject: [TASK NINJA] Updates for " + request.IssueType + ": " + request.Title + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	request.Timestamp = time.Now()
	body, err := ParseTemplate("./service/template.html", &request)
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
	fmt.Println("Email sent successfully!")
}

func ParseTemplate(fileName string, request interface{}) (string, error) {
	t, err := template.ParseFiles("./service/template.html")
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
