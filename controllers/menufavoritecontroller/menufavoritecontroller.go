package menufavoritecontroller

import (
	"eatdah/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var menufavorite []models.MenuFavorite

	models.DB.Find(&menufavorite)
	c.JSON(http.StatusOK, gin.H{"menu_favorite": menufavorite})
}

func Show(c *gin.Context) {

	var menufavorite models.MenuFavorite

	id := c.Param("id")

	if err := models.DB.First(&menufavorite, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"MenuFavorite": menufavorite})

}

func Create(c *gin.Context) {

	var menufavorite models.MenuFavorite

	if err := c.ShouldBindJSON(&menufavorite); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&menufavorite)
	c.JSON(http.StatusOK, gin.H{"MenuFavorite": menufavorite})

}

func Update(c *gin.Context) {

	var menufavorite models.MenuFavorite
	id := c.Param("id")

	if err := c.ShouldBindJSON(&menufavorite); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&menufavorite).Where("id_menufavorite = ?", id).Updates(&menufavorite).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Mengupdate Product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diperbarui"})

}

func Delete(c *gin.Context) {

	var menufavorite models.MenuFavorite

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&menufavorite, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Menghapus Product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})

}
