package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "ID of the task to delete"
// @Success 200 {string} string "Task deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/delete [delete]
func DeleteTask(c *gin.Context) {
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
	//extract the user id associated with the current session
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		taskId := c.Query("id")
		rows, err := db.Exec("delete from tasks where user_id=? and id=?", userId, taskId)
		if err != nil {
			utils.Logger.Err(err).Msg("Error occurred while executing query")
			c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		}
		result, err := rows.RowsAffected()
		if err != nil {
			utils.Logger.Err(err).Msg("Error occurred while getting the number of affected rows")
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if result == 0 {
			c.JSON(http.StatusNotFound, "no tasks found for the user")
			return
		}

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	utils.Logger.Info().Msg("Task deleted successfully!")
	c.JSON(http.StatusOK, "Task Deleted !!!")
}
