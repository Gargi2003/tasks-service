package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// lists down all tasks associated with the logged-in user
func ListTasks(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SecretKey), nil
	})
	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var tasks []utils.Task
	// Extract the user ID associated with the current session
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]

		rows, err := db.Query("SELECT id, title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points,reporter, comments, status FROM tasks WHERE user_id=?", userId)
		if err != nil {
			utils.Logger.Err(err).Msg("Error occurred while executing query")
			c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
			return
		}

		for rows.Next() {
			task := utils.Task{}
			var createdAt, updatedAt string
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &createdAt, &updatedAt, &task.UserID, &task.IssueType, &task.Assignee, &task.Sprint, &task.ProjectId, &task.StoryPoints, &task.Reporter, &task.Comments, &task.Status)
			if err != nil {
				utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
			}

			// Parse the string values into time.Time
			task.CreatedAt, _ = time.Parse("2006-01-02", createdAt)
			task.UpdatedAt, _ = time.Parse("2006-01-02", updatedAt)

			tasks = append(tasks, task)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, tasks)
}
