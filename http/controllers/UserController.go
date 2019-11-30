package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/fatih/structs"
)

type UserController struct {
	Controller
}

func (UserController) List(c *gin.Context) {
	users, err := services.UserService{}.FindAll()
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, users)
}

func (UserController) Get(c *gin.Context) {
	userOid := GetUserID(c)

	user, err := services.UserService{}.FindById(userOid)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, user)
}

func (UserController) Patch(c *gin.Context) {
	userOid := GetUserID(c)

	var userPatch models.UserPatch
	err := c.ShouldBind(&userPatch)
	if utils.HttpError(c, 400, err) { return }

	patchMap := structs.Map(userPatch)

	if PatchValid(patchMap, models.UserPatchSchema) {
		err = services.UserService{}.Update(userOid, &patchMap)
		if utils.HttpError(c, 500, err) { return }
	} else {
		if utils.HttpError(c, 400, errors.New("invalid patch")) { return }
	}

}

func (UserController) Delete(c *gin.Context) {
	userOid := GetUserID(c)

	err := services.UserService{}.Delete(userOid)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, err)
}
