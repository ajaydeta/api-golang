package Controllers

import (
	"belajar_golang/Config"
	"belajar_golang/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var taks []Models.Taks
var taksKategori []Models.KategoriTaks

// GetTaks untuk ambil semua taks
func GetTaks(c *gin.Context) {
	type ResultJoin struct {
		ID         int       `json:"id"`
		Nama       string    `json:"nama"`
		DueDate    time.Time `json:"due_date"`
		Deskripsi  string    `json:"deskripsi"`
		IDKategori int       `json:"id_kategori"`
		Kategori   string    `json:"kategori"`
	}

	idKategori := c.Query("idkategori")
	if len(idKategori) != 0 {

		result := Config.DB.Table("taks").
			Select("taks.*, kategori_taks.nama as kategori").
			Joins("join kategori_taks on taks.id_kategori = kategori_taks.id").
			Where("taks.id_kategori = ?", idKategori).
			//Scan(&ResultJoin{})
			Find(&taks)

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id kategori not found"})
			return
		}

		res := gin.H{
			"data":  result.Value,
			"error": result.Error,
		}
		c.JSON(http.StatusOK, res)
		println(len(idKategori))
		return
	}

	result := Config.DB.Find(&taks)
	res := gin.H{
		"data":  result.Value,
		"error": result.Error,
	}

	c.JSON(http.StatusOK, res)
}
func FindTaks(c *gin.Context) {
	result := Config.DB.Find(&taks, "id = ?", c.Param("id"))
	if result.RowsAffected != 0 {
		res := gin.H{
			"data":  result.Value,
			"error": result.Error,
		}
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Request id not found"})
}

type CreateTaksInput struct {
	ID         int       `json:"id"`
	Nama       string    `json:"nama"`
	DueDate    time.Time `json:"due_date"`
	Deskripsi  string    `json:"deskripsi"`
	IDKategori int       `json:"id_kategori"`
}

func CreateTaks(c *gin.Context) {
	var input CreateTaksInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taks := Models.Taks{
		ID:         input.ID,
		Nama:       input.Nama,
		DueDate:    input.DueDate,
		Deskripsi:  input.Deskripsi,
		IDKategori: input.IDKategori,
	}

	result := Config.DB.Create(&taks)
	res := gin.H{
		"data":  result.Value,
		"error": result.Error,
	}

	c.JSON(http.StatusOK, res)
}

type UpdateTaksInput struct {
	Nama       string    `json:"nama"`
	DueDate    time.Time `json:"due_date"`
	Deskripsi  string    `json:"deskripsi"`
	IDKategori int       `json:"id_kategori"`
}

func UpadateTaks(c *gin.Context) {
	result := Config.DB.First(&taks, "id = ?", c.Param("id"))
	if result.RowsAffected == 0 {
		res := gin.H{
			"error": "Request id not found",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var input UpdateTaksInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorss": err.Error()})
		return
	}

	Config.DB.Model(&taks).Where("id = ?", c.Param("id")).Update(&input)

	c.JSON(http.StatusOK, gin.H{"data": &input})
}

func DeleteTaks(c *gin.Context) {
	result := Config.DB.First(&taks, "id = ?", c.Param("id"))
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Request id not found"})
		return
	}

	Config.DB.Delete(&taks, "id =  ?", c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": true})
}