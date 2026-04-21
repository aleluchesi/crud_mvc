package controller

import (
	"fmt"
	"net/http"

	logger "github.com/aleluchesi/crud_mvc/src/configuration"
	"github.com/aleluchesi/crud_mvc/src/configuration/validation"
	"github.com/aleluchesi/crud_mvc/src/model/request"
	"github.com/aleluchesi/crud_mvc/src/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error typing to marshal object", err, zap.String("journey", "createUser"))
		restErr := validation.ValodateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		Id:    "1",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, response)

}
