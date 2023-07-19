package menucontroller

import (
	"eatdah/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var menu []models.Menu

	models.DB.Find(&menu)
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func Show(c *gin.Context) {
	var menu models.Menu

	id := c.Param("id")

	if err := models.DB.First(&menu, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})

}

func Create(c *gin.Context) {

	var menu models.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&menu)
	c.JSON(http.StatusOK, gin.H{"menu": menu})

}

func Update(c *gin.Context) {

	var menu models.Menu
	id := c.Param("id")

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&menu).Where("id_menu = ?", id).Updates(&menu).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Mengupdate Product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diperbarui"})

}

func Delete(c *gin.Context) {

	var menu models.Menu

	var input struct {
		Id_menu json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id_menu.Int64()
	if models.DB.Delete(&menu, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Menghapus Product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})

}
