package controllers

import (
	"App/database/models"
	"App/utils"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	Controller
}

func (controller CourseController) List(c *gin.Context) {
	courses := models.CourseRepository{}.FindAll()
	c.JSON(200, courses)
}

func (controller CourseController) Post(c *gin.Context) {
	var course models.Course
	err := c.Bind(&course)
	utils.HttpError(c, 400, err)
	course.Store()
	utils.HttpError(c, 500, err)
	c.JSON(200, course)
}

func (controller CourseController) Patch(c *gin.Context) {
	if value, ok := c.Get("course"); ok {
		course := value.(models.Course)
		var coursePatch models.CoursePatch
		err := c.Bind(&coursePatch)
		utils.HttpError(c, 400, err)
		course.ApplyPatch(coursePatch)
		course.Store()
		c.JSON(202, course)
	}

}

func (controller CourseController) Get(c *gin.Context) {
	if value, ok := c.Get("course"); ok {
		course := value.(models.Course)
		c.JSON(200, course)
	}
}

func (controller CourseController) Delete(c *gin.Context) {
	if value, ok := c.Get("course"); ok {
		course := value.(models.Course)
		err := course.Drop()
		utils.HttpError(c, 500, err)
		c.JSON(204, course)
	}
}
