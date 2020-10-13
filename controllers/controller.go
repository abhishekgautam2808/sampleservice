package controllers

import (
	"net/http"
	"strconv"

	"github.com/abhishekgautam2808/sampleservice/models"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var users []models.User
	err := models.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetSingleUser(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = uint(id)
	err = models.DB.Where("id = ?", user.ID).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
