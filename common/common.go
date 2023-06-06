package common

import (
	"database/sql"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

type PreviousResponse struct {
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
	Status      string    `json:"status"`
}
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"`
	IssueType   string    `json:"issue_type"`
	Assignee    string    `json:"assignee"`
	Sprint      int       `json:"sprint_id"`
	ProjectId   int       `json:"project_id"`
	StoryPoints int       `json:"points"`
	Reporter    string    `json:"reporter"`
	Comments    string    `json:"comments"`
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
