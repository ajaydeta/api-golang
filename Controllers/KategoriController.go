package Controllers

import (
	"belajar_golang/Config"
	"belajar_golang/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var kategori []Models.KategoriTaks

type InputKategori struct {
	Nama string `json:"nama"`
}

func GetKategori(c *gin.Context) {
	result := Config.DB.Find(&kategori)
	res := gin.H{
		"data":  result.Value,
		"error": result.Error,
	}
	c.JSON(http.StatusOK, res)
}

func CreateKategori(c *gin.Context) {
	var input InputKategori
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kategori := Models.KategoriTaks{
		Nama: input.Nama,
	}

	result := Config.DB.Create(&kategori)

	res := gin.H{
		"data":  result.Value,
		"error": result.Error,
	}
	c.JSON(http.StatusOK, res)
}

func UpdateKategori(c *gin.Context) {
	var input InputKategori

	row := Config.DB.First(&kategori, "id = ?", c.Param("id")).RowsAffected
	if row == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := Config.DB.Model(&kategori).Where("id = ?", c.Param("id")).Update(&input)

	res := gin.H{
		"data":  true,
		"error": result.Error,
	}
	c.JSON(http.StatusOK, res)
}

func DeleteKategori(c *gin.Context) {
	row := Config.DB.First(&kategori, "id = ?", c.Param("id")).RowsAffected
	if row == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
		return
	}

	Config.DB.Delete(&kategori, "id = ?", c.Param("id"))
	res := gin.H{
		"data": true,
	}
	c.JSON(http.StatusOK, res)
}
