package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type EditRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
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

		if req.Completed {
			updateQuery += " completed=true,"
		} else {
			updateQuery += " completed=false,"
		}

		// Remove the trailing comma from the query
		updateQuery = updateQuery[:len(updateQuery)-1]

		// Add the WHERE condition
		updateQuery += " WHERE user_id=? and id=?"

		taskId := c.Query("id")
		// Add the user ID parameter to the query parameters
		queryParams = append(queryParams, userId, taskId)

		// Execute the update query
		_, err := db.Query(updateQuery, queryParams...)
		if err != nil {
			utils.Logger.Err(err).Msg("Error executing update query")
			c.JSON(http.StatusInternalServerError, "Error updating tasks")
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// Handle the response
	c.JSON(http.StatusOK, "Tasks updated successfully")
}
