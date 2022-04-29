package controllers

import (
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm/clause"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("UserData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == AppJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("UserData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo := []models.Photo{}
	response := []map[string]interface{}{}
	_ = response
	err := db.Preload("User").Where("user_id=?", userID).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range Photo {
		nestedData := map[string]interface{}{
			"email":    Photo[i].User.Email,
			"username": Photo[i].User.Username,
		}
		data := map[string]interface{}{
			"id":         Photo[i].ID,
			"title":      Photo[i].Title,
			"caption":    Photo[i].Caption,
			"photo_url":  Photo[i].PhotoUrl,
			"user_id":    Photo[i].UserID,
			"created_at": Photo[i].CreatedAt,
			"updated_at": Photo[i].UpdateAt,
			"User":       nestedData,
		}

		response = append(response, data)
	}

	c.JSON(http.StatusOK, response)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoID"))

	if contentType == AppJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Photo).Where("id=?", photoID).Clauses(clause.Returning{}).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdateAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(c.Param("photoID"))

	var Photo models.Photo

	err := db.Model(&Photo).Delete(&Photo, photoID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been deleted",
	})
}
