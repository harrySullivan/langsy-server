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

type CourseController struct {
	Controller
}

func (CourseController) List(c *gin.Context) {
	courses, err := services.CourseService{}.FindAll()
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, courses)
}

func (CourseController) Get(c *gin.Context) {
	courseId := c.Param("courseId")
	courseOid, err := primitive.ObjectIDFromHex(courseId)
	if utils.HttpError(c, 400, err) { return }

	course, err := services.CourseService{}.FindById(&courseOid)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, course)
}

func (CourseController) Post(c *gin.Context) {
	var course models.Course
	err := c.Bind(&course)
	if utils.HttpError(c, 400, err) { return }

	iCourse, iErr := services.CourseService{}.Insert(&course)
	if utils.HttpError(c, 500, iErr) { return }

	c.JSON(200, iCourse)
}

func (CourseController) Patch(c *gin.Context) {
	courseId := c.Param("courseId")
	courseOid, err := primitive.ObjectIDFromHex(courseId)
	if utils.HttpError(c, 400, err) { return }

	var coursePatch models.CoursePatch
	err = c.ShouldBind(&coursePatch)
	if utils.HttpError(c, 400, err) { return }

	patchMap := structs.Map(coursePatch)

	if PatchValid(patchMap, models.CoursePatchSchema) {
		err = services.CourseService{}.Update(&courseOid, &patchMap)
		if utils.HttpError(c, 500, err) { return }
	} else {
		if utils.HttpError(c, 400, errors.New("invalid patch")) { return }
	}

}

func (CourseController) Delete(c *gin.Context) {
	courseId := c.Param("courseId")
	courseOid, err := primitive.ObjectIDFromHex(courseId)
	utils.HttpError(c, 400, err)

	err = services.CourseService{}.Delete(&courseOid)
	utils.HttpError(c, 500, err)

	c.JSON(200, err)
}
