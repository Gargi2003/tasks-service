package handlers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	utils "tasks/common"
	"tasks/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateTasks godoc
// @Summary Create a new task
// @Description Create a new task for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Param title formData string true "Title of the task"
// @Param description formData string false "Description of the task"
// @Param issue_type formData string true "Type of the task (e.g., bug, feature)"
// @Param assignee formData string true "Assignee of the task"
// @Param sprint formData int false "ID of the sprint the task belongs to"
// @Param project_id formData int false "ID of the project the task belongs to"
// @Param points formData int false "Points/effort estimation for the task"
// @Param reporter formData string true "Reporter of the task"
// @Param comments formData string false "Additional comments for the task"
// @Param status formData string false "Status of the task"
// @Success 200 {string} string "Task created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks [post]
func CreateTasks(c *gin.Context) {

	//get the jwt token from the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		utils.Logger.Err(err).Msg("Error getting token")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//parse the token and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SecretKey), nil
	})

	//bind the request to the struct
	var req utils.Task
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusInternalServerError, "Error binding req object")
	}

	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	//extract the user id associated with the current session
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		url := fmt.Sprint("http://localhost:8080/users/get?id=", userId)
		response, err := http.Get(url)
		if err != nil {
			utils.Logger.Err(err).Msg("Error calling users API")
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			utils.Logger.Err(err).Msg("Error reading response body")
		}
		username := string(body)
		//introduce check to see if user creating the task
		//is creating the task for him or someone else
		//if someone else then send email to the assignee
		if username != req.Assignee {
			// Send email to the assignee
			service.SendEmailForCreatedIssue(req)
		}
		//create the tasks
		var result sql.Result
		if result, err = db.Exec("insert into tasks (title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points, reporter, comments, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", req.Title, req.Description, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), userId, req.IssueType, req.Assignee, req.Sprint, req.ProjectID, req.Points, req.Reporter, req.Comments, req.Status); err != nil {
			utils.Logger.Err(err).Msg("Error inserting data into tasks table")
			c.JSON(http.StatusInternalServerError, "Error inserting data into tasks table")
			return
		}
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			utils.Logger.Err(err).Msg("Error retrieving last inserted ID")
			c.JSON(http.StatusInternalServerError, "Error retrieving last inserted ID")
			return
		}
		req.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		req.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		taskId := fmt.Sprintf("%d", lastInsertId)
		if err := utils.UpdateAudit(nil, &req, userId, taskId); err != nil {
			utils.Logger.Err(err).Msg("Error updating audit table")
			c.JSON(http.StatusInternalServerError, "Error updating audit table")
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, "Task created successfully!")
}
