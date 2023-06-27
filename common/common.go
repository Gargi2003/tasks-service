package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

type Task struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Count          int
	PreviousFields map[string]interface{}
	Timestamp      time.Time
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UserID         int       `json:"user_id"`
	IssueType      string    `json:"issue_type"`
	Assignee       string    `json:"assignee"`
	Sprint         int       `json:"sprint_id"`
	ProjectID      int       `json:"project_id"`
	Points         int       `json:"points"`
	Reporter       string    `json:"reporter"`
	Comments       string    `json:"comments"`
	Status         string    `json:"status"`
}
type Tasks_Audit struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
	IssueType   string    `json:"issue_type"`
	Assignee    string    `json:"assignee"`
	Sprint      int       `json:"sprint_id"`
	ProjectID   int       `json:"project_id"`
	Points      int       `json:"points"`
	Reporter    string    `json:"reporter"`
	Comments    string    `json:"comments"`
	TaskID      string    `json:"task_id"`
	Status      string    `json:"status"`
}

const (
	[REDACTED_USERNAME]
	[REDACTED_PASSWORD]
	Dbname     = "todo_manager"
	Topology   = "tcp"
	Port       = "localhost:3306"
	DriverName = "mysql"
	SecretKey  = "khsiudjsb12jhb4!"
)

var Logger zerolog.Logger = zerolog.New(os.Stdout)

func DBConn(user string, password string, dbname string, port string) (*sql.DB, error) {
	dataSourceName := ConstructURL(user, password, dbname, port)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		Logger.Err(err).Msg("Error connecting to database")
		return nil, err
	}
	return db, nil
}

func ConstructURL(user string, password string, dbname string, port string) string {
	var sb strings.Builder
	sb.WriteString(user)
	sb.WriteString(":")
	sb.WriteString(password)
	sb.WriteString("@")
	sb.WriteString(Topology)
	sb.WriteString("(")
	sb.WriteString(port)
	sb.WriteString(")")
	sb.WriteString("/")
	sb.WriteString(dbname)

	return sb.String()
}

func GetTaskByID(c *gin.Context, taskID string) (*Task, error) {

	//get the jwt token from the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		Logger.Err(err).Msg("Error getting token")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//parse the token and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		apiURL := fmt.Sprintf("http://localhost:8081/tasks/get?userid=%s&taskid=%s", userId, taskID)

		response, err := http.Get(apiURL)
		if err != nil {
			Logger.Err(err).Msg("Error calling users API")
		}
		defer response.Body.Close()

		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API request failed with status: %d", response.StatusCode)
		}

		var task Task
		err = json.NewDecoder(response.Body).Decode(&task)
		if err != nil {
			return nil, err
		}
		return &task, nil
	}

	return nil, nil
}

func UpdateAudit(previousTask *Task, req *Task, userId interface{}, taskId interface{}) error {
	//connect to db
	db, err := DBConn(Username, Password, Dbname, Port)
	if err != nil {
		Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	if previousTask != nil {
		if req.Title == "" {
			req.Title = previousTask.Title
		}
		if req.Description == "" {
			req.Description = previousTask.Description
		}
		if req.IssueType == "" {
			fmt.Println("issuetype empty")
			req.IssueType = previousTask.IssueType
			fmt.Println("issuetype", req.IssueType)
		}

		if req.Assignee == "" {
			req.Assignee = previousTask.Assignee
		}

		if req.Sprint == 0 {
			req.Sprint = previousTask.Sprint
		}
		if req.ProjectID == 0 {
			req.ProjectID = previousTask.ProjectID
		}

		if req.Points == 0 {
			req.Points = previousTask.Points
		}

		if req.Reporter == "" {
			req.Reporter = previousTask.Reporter
		}

		if req.Comments == "" {
			req.Comments = previousTask.Comments
		}

		if req.Status == "" {
			req.Status = previousTask.Status
		}
	}

	//create the tasks
	_, err1 := db.Query("insert into tasks_audit (title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points, reporter, comments, status, task_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", req.Title, req.Description, req.CreatedAt, req.UpdatedAt, userId, req.IssueType, req.Assignee, req.Sprint, req.ProjectID, req.Points, req.Reporter, req.Comments, req.Status, taskId)

	if err1 != nil {
		return err1
	}

	return nil
}
