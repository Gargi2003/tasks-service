package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
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
		//create the tasks
		db.Query("insert into tasks (title, description, completed, created_at, updated_at, user_id) VALUES (?, ?, ?, ?, ?, ?)", req.Title, req.Description, req.Completed, time.Now(), time.Now(), userId)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, "Task created successfully!")
}
