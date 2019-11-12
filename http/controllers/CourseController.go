package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/utils"

	"github.com/gin-gonic/gin"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseController struct {
	Controller
}

func (CourseController) List(c *gin.Context) {
	courses, err := services.CourseService{}.FindAll()
	utils.HttpError(c, 500, err)

	c.JSON(200, courses)
}

func (CourseController) Get(c *gin.Context) {
	courseId := c.Param("courseId")
	courseOid, err := primitive.ObjectIDFromHex(courseId)
	utils.HttpError(c, 400, err)

	course, err := services.CourseService{}.FindById(&courseOid)
	utils.HttpError(c, 500, err)

	c.JSON(200, course)
}

func (CourseController) Post(c *gin.Context) {
	var course models.Course
	err := c.Bind(&course)
	utils.HttpError(c, 400, err)

	iCourse, iErr := services.CourseService{}.Insert(&course)
	utils.HttpError(c, 500, iErr)

	c.JSON(200, iCourse)
}

func (CourseController) Patch(c *gin.Context) {
	courseId := c.Param("courseId")
	courseOid, err := primitive.ObjectIDFromHex(courseId)
	utils.HttpError(c, 400, err)

	var coursePatch models.CoursePatch
	err = c.ShouldBind(&coursePatch)
	utils.HttpError(c, 400, err)

	patchMap := structs.Map(coursePatch)

	err = services.CourseService{}.Update(&courseOid, &patchMap)
	utils.HttpError(c, 500, err)
}

func (CourseController) Delete(c *gin.Context) {
	courseId := c.Param("courseId")
	courseOid, err := primitive.ObjectIDFromHex(courseId)
	utils.HttpError(c, 400, err)

	err = services.CourseService{}.Delete(&courseOid)
	utils.HttpError(c, 500, err)

	c.JSON(200, err)
}
