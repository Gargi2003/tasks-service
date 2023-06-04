package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	utils "tasks/common"
	service "tasks/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

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
	var req CreateRequest
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
		url := fmt.Sprint("http://localhost:8080/get?id=", userId)
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
			service.SendEmail()
		}
		//create the tasks
		db.Query("insert into tasks (title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points, reporter, comments, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", req.Title, req.Description, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), userId, req.IssueType, req.Assignee, req.Sprint, req.ProjectId, req.StoryPoints, req.Reporter, req.Comments, req.Status)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, "Task created successfully!")
}
