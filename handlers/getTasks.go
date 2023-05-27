package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// lists down all tasks associated with the loggedin user
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
	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var tasks []utils.Task
	//extract the user id associated with the current session
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]

		rows, err := db.Query("select id,title,description,completed,created_at,updated_at from tasks where user_id=?", userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		}
		for rows.Next() {
			var task utils.Task
			rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
			tasks = append(tasks, task)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, tasks)
}
