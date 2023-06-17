package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"
	"tasks/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// EditTasks updates a task with the provided details.
// @Summary Edit Task
// @Description Update a task with the provided details
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Task ID"
// @Success 200 {string} string "Tasks updated successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/edit [put]
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
	var req utils.Task
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusInternalServerError, "Error binding req object")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		taskId := c.Query("id")

		// Fetch the task before the update
		previousTask, err := utils.GetTaskByID(c, taskId)
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
		if req.ProjectID != 0 {
			updateQuery += " project_id=?,"
			queryParams = append(queryParams, req.ProjectID)
		}

		if req.Points != 0 {
			updateQuery += " points=?,"
			queryParams = append(queryParams, req.Points)
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

		result, err := utils.GetTaskByID(c, taskId)
		req.CreatedAt = result.CreatedAt
		req.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		if err := utils.UpdateAudit(previousTask, &req, userId, taskId); err != nil {
			utils.Logger.Err(err).Msg("Error updating audit table")
			c.JSON(http.StatusInternalServerError, "Error updating audit table")
			return
		}

		// Fetch the updated task
		updatedTask, err := utils.GetTaskByID(c, taskId)
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
