package controller

import (
	"fmt"
	"log"

	"github.com/aleluchesi/crud_mvc/src/configuration/validation"
	"github.com/aleluchesi/crud_mvc/src/model/request"
	"github.com/aleluchesi/crud_mvc/src/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error %s", err.Error())
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

	c.JSON(200, response)

}
