package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	Controller
}

func (UserController) List(c *gin.Context) {
	users, err := services.UserService{}.FindAll()
	utils.HttpError(c, 500, err)

	c.JSON(200, users)
}

func (UserController) Get(c *gin.Context) {
	userId := c.Param("userId")
	userOid, err := primitive.ObjectIDFromHex(userId)
	utils.HttpError(c, 400, err)

	user, err := services.UserService{}.FindById(&userOid)
	utils.HttpError(c, 500, err)

	c.JSON(200, user)
}

func (UserController) Post(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	utils.HttpError(c, 400, err)

	iCourse, iErr := services.UserService{}.Insert(&user)
	utils.HttpError(c, 500, iErr)

	c.JSON(200, iCourse)
}

func (UserController) Patch(c *gin.Context) {
	userId := c.Param("userId")
	userOid, err := primitive.ObjectIDFromHex(userId)
	utils.HttpError(c, 400, err)

	var userPatch models.UserPatch
	err = c.ShouldBind(&userPatch)
	utils.HttpError(c, 400, err)

	patchMap := structs.Map(userPatch)

	if PatchValid(patchMap, models.UserPatchSchema) {
		err = services.UserService{}.Update(&userOid, &patchMap)
		utils.HttpError(c, 500, err)
	} else {
		utils.HttpError(c, 400, errors.New("invalid patch"))
	}

}

func (UserController) Delete(c *gin.Context) {
	userId := c.Param("userId")
	userOid, err := primitive.ObjectIDFromHex(userId)
	utils.HttpError(c, 400, err)

	err = services.UserService{}.Delete(&userOid)
	utils.HttpError(c, 500, err)

	c.JSON(200, err)
}
