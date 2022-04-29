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

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("UserData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == AppJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMedia.UserID = userID

	err := db.Debug().Create(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}
func GetsocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("UserData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMedia := []models.SocialMedia{}
	response := []map[string]interface{}{}
	_ = response
	err := db.Preload("User").Where("user_id=?", userID).Find(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range socialMedia {
		nestedData := map[string]interface{}{
			"id":       socialMedia[i].User.ID,
			"username": socialMedia[i].User.Username,
		}
		data := map[string]interface{}{
			"id":               socialMedia[i].ID,
			"name":             socialMedia[i].Name,
			"social_media_url": socialMedia[i].SocialMediaUrl,
			"user_id":          socialMedia[i].UserID,
			"created_at":       socialMedia[i].CreatedAt,
			"updated_at":       socialMedia[i].UpdateAt,
			"User":             nestedData,
		}

		response = append(response, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": response})
}

func UpdatesocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	socialMedia := models.SocialMedia{}

	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))

	if contentType == AppJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	err := db.Model(&socialMedia).Where("id=?", socialMediaID).Clauses(clause.Returning{}).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaUrl: socialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdateAt,
	})
}

func DeletesocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))

	var socialMedia models.SocialMedia

	err := db.Model(&socialMedia).Delete(&socialMedia, socialMediaID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been sucessfully deleted",
	})
}
