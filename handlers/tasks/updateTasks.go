package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	utils "tasks/common"
	"tasks/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type EditRequest struct {
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

func EditTasks(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		// Handle the error
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SecretKey), nil
	})
	// Connect to the database
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	// Bind the request to the struct
	var req EditRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusInternalServerError, "Error binding req object")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		taskId := c.Query("id")

		// Fetch the task before the update
		previousTask, err := GetTaskByID(c, taskId)
		if err != nil {
			utils.Logger.Err(err).Msg("Error fetching previous task")
			c.JSON(http.StatusInternalServerError, "Error fetching previous task")
			return
		}
		// Build the update query dynamically based on non-empty fields
		updateQuery := "UPDATE tasks SET"
		var queryParams []interface{}

		if req.Title != "" {
			updateQuery += " title=?,"
			queryParams = append(queryParams, req.Title)
		}

		if req.Description != "" {
			updateQuery += " description=?,"
			queryParams = append(queryParams, req.Description)
		}
		if req.IssueType != "" {
			updateQuery += " issue_type=?,"
			queryParams = append(queryParams, req.IssueType)
		}

		if req.Assignee != "" {
			updateQuery += " assignee=?,"
			queryParams = append(queryParams, req.Assignee)
		}

		if req.Sprint != 0 {
			updateQuery += " sprint_id=?,"
			queryParams = append(queryParams, req.Sprint)
		}
		if req.ProjectId != 0 {
			updateQuery += " project_id=?,"
			queryParams = append(queryParams, req.ProjectId)
		}

		if req.StoryPoints != 0 {
			updateQuery += " points=?,"
			queryParams = append(queryParams, req.StoryPoints)
		}

		if req.Reporter != "" {
			updateQuery += " reporter=?,"
			queryParams = append(queryParams, req.Reporter)
		}

		if req.Comments != "" {
			updateQuery += " comments=?,"
			queryParams = append(queryParams, req.Comments)
		}

		if req.Status != "" {
			updateQuery += " status=?,"
			queryParams = append(queryParams, req.Status)
		}

		// Add the updated_at field
		updateQuery += " updated_at=? "

		// Format the current time in the desired format
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		// Append the updated_at parameter to the query parameters
		queryParams = append(queryParams, currentTime)
		// Remove the trailing comma from the query
		updateQuery = updateQuery[:len(updateQuery)-1]

		// Add the WHERE condition
		updateQuery += " WHERE user_id=? and id=?"

		utils.Logger.Info().Msg(updateQuery)

		// Add the user ID parameter to the query parameters
		queryParams = append(queryParams, userId, taskId)

		// Execute the update query
		_, err1 := db.Query(updateQuery, queryParams...)
		if err1 != nil {
			utils.Logger.Err(err1).Msg("Error executing update query")
			c.JSON(http.StatusInternalServerError, "Error updating tasks")
			return
		}

		// Fetch the updated task
		updatedTask, err := GetTaskByID(c, taskId)
		if err != nil {
			utils.Logger.Err(err).Msg("Error fetching updated task")
			c.JSON(http.StatusInternalServerError, "Error fetching updated task")
			return
		}
		// Send email with updated and previous tasks
		service.SendEmailForUpdatedIssue(previousTask, updatedTask)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// Handle the response
	c.JSON(http.StatusOK, "Tasks updated successfully")
}

func GetTaskByID(c *gin.Context, taskID string) (*utils.PreviousResponse, error) {

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
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		apiURL := fmt.Sprintf("http://localhost:8081/tasks/get?userid=%s&taskid=%s", userId, taskID)

		response, err := http.Get(apiURL)
		if err != nil {
			utils.Logger.Err(err).Msg("Error calling users API")
		}
		defer response.Body.Close()

		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API request failed with status: %d", response.StatusCode)
		}

		var task utils.PreviousResponse
		err = json.NewDecoder(response.Body).Decode(&task)
		if err != nil {
			return nil, err
		}
		return &task, nil
	}

	return nil, nil
}
