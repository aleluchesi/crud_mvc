package controller

import (
	"fmt"

	"github.com/aleluchesi/crud_mvc/src/configuration/rest_err"
	"github.com/aleluchesi/crud_mvc/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRedquestError(
			fmt.Sprintf("There are some incorrect fields, error %s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(UserRequest)

}
